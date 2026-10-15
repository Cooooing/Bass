package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"platform/internal/biz/model"
	bizrepo "platform/internal/biz/repo"
	"platform/internal/enum"
)

var _ bizrepo.HolidayClient = (*HolidayClient)(nil)

type HolidayClient struct {
	client    *http.Client
	baseURL   string
	userAgent string
}

func NewHolidayClient(client *http.Client) bizrepo.HolidayClient {
	return &HolidayClient{
		client:    client,
		baseURL:   "https://timor.tech",
		userAgent: "Bass/1.0 (+https://fishpi.cn)",
	}
}

func (c *HolidayClient) Fetch(ctx context.Context, date time.Time) (*model.HolidayInfo, error) {
	day := date.Format("2006-01-02")
	info, err := c.fetchInfo(ctx, day)
	if err != nil {
		return nil, err
	}

	switch info.Type {
	case enum.HolidayTypeWeekend, enum.HolidayTypeFestival:
		rest, err := c.fetchWorkdayRest(ctx, day)
		if err != nil {
			return nil, err
		}
		info.DaysRemaining = rest
	case enum.HolidayTypeWorkday, enum.HolidayTypeMakeupWorkday:
		name, rest, err := c.fetchNextHoliday(ctx, day)
		if err != nil {
			return nil, err
		}
		info.HolidayName = name
		info.DaysToHoliday = rest
	}
	return info, nil
}

func (c *HolidayClient) fetchInfo(ctx context.Context, day string) (*model.HolidayInfo, error) {
	var resp timorInfoResp
	if err := c.getJSON(ctx, "/api/holiday/info/"+day, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("holiday info error code %d", resp.Code)
	}
	holidayType, err := c.holidayTypeFromCode(resp.Type.Type)
	if err != nil {
		return nil, err
	}
	return &model.HolidayInfo{
		Date:    c.parseHolidayDate(day),
		Type:    holidayType,
		DayName: resp.Type.Name,
	}, nil
}

func (c *HolidayClient) fetchWorkdayRest(ctx context.Context, day string) (int32, error) {
	var resp timorWorkdayResp
	if err := c.getJSON(ctx, "/api/holiday/workday/next/"+day, &resp); err != nil {
		return 0, err
	}
	if resp.Code != 0 {
		return 0, fmt.Errorf("holiday workday error code %d", resp.Code)
	}
	return resp.Workday.Rest, nil
}

func (c *HolidayClient) fetchNextHoliday(ctx context.Context, day string) (string, int32, error) {
	var resp timorNextResp
	if err := c.getJSON(ctx, "/api/holiday/next/"+day+"?type=Y&week=Y", &resp); err != nil {
		return "", 0, err
	}
	if resp.Code != 0 {
		return "", 0, fmt.Errorf("holiday next error code %d", resp.Code)
	}
	return resp.Holiday.Name, resp.Holiday.Rest, nil
}

func (c *HolidayClient) getJSON(ctx context.Context, path string, target any) (err error) {
	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(requestCtx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return fmt.Errorf("create holiday request %s: %w", path, err)
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("request holiday %s: %w", path, err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); err == nil && closeErr != nil {
			err = fmt.Errorf("close holiday response %s: %w", path, closeErr)
		}
	}()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request holiday %s: status %d", path, resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decode holiday %s: %w", path, err)
	}
	return nil
}

func (c *HolidayClient) holidayTypeFromCode(code int) (enum.HolidayType, error) {
	switch code {
	case 0:
		return enum.HolidayTypeWorkday, nil
	case 1:
		return enum.HolidayTypeWeekend, nil
	case 2:
		return enum.HolidayTypeFestival, nil
	case 3:
		return enum.HolidayTypeMakeupWorkday, nil
	default:
		return "", fmt.Errorf("unknown holiday type %d", code)
	}
}

func (c *HolidayClient) parseHolidayDate(day string) time.Time {
	parsed, err := time.ParseInLocation("2006-01-02", day, time.FixedZone("Asia/Shanghai", 8*60*60))
	if err != nil {
		return time.Time{}
	}
	return parsed
}

type timorInfoResp struct {
	Code int `json:"code"`
	Type struct {
		Type int    `json:"type"`
		Name string `json:"name"`
	} `json:"type"`
}

type timorWorkdayResp struct {
	Code    int `json:"code"`
	Workday struct {
		Rest int32 `json:"rest"`
	} `json:"workday"`
}

type timorNextResp struct {
	Code    int `json:"code"`
	Holiday struct {
		Name string `json:"name"`
		Rest int32  `json:"rest"`
	} `json:"holiday"`
}
