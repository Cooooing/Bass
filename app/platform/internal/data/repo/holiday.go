package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"platform/internal/biz/model"
	bizrepo "platform/internal/biz/repo"
)

var _ bizrepo.HolidayClient = (*HolidayClient)(nil)

type HolidayClient struct {
	client    *http.Client
	baseURL   string
	userAgent string
	location  *time.Location
}

func NewHolidayClient(client *http.Client) bizrepo.HolidayClient {
	return &HolidayClient{
		client:  client,
		baseURL: "https://timor.tech",
		// timor.tech 拒绝默认的 Go-http-client User-Agent，必须显式携带一个普通 UA。
		userAgent: "Bass/1.0 (+https://fishpi.cn)",
		location:  time.FixedZone("Asia/Shanghai", 8*60*60),
	}
}

// FetchYear 一次性读取整年的节日与调休配置，后续按自然日展开都在本地完成。
func (c *HolidayClient) FetchYear(ctx context.Context, year int) (map[string]*model.HolidayEntry, error) {
	var resp timorYearResp
	if err := c.getJSON(ctx, fmt.Sprintf("/api/holiday/year/%d/", year), &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("holiday year error code %d", resp.Code)
	}
	entries := make(map[string]*model.HolidayEntry, len(resp.Holiday))
	for _, raw := range resp.Holiday {
		if raw == nil || raw.Date == "" {
			continue
		}
		date, err := c.parseDate(raw.Date)
		if err != nil {
			return nil, err
		}
		entries[raw.Date] = &model.HolidayEntry{
			Date:    date,
			Name:    raw.Name,
			Holiday: raw.Holiday,
		}
	}
	return entries, nil
}

func (c *HolidayClient) parseDate(value string) (time.Time, error) {
	date, err := time.ParseInLocation("2006-01-02", value, c.location)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse holiday date %s: %w", value, err)
	}
	return date, nil
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

type timorYearResp struct {
	Code    int                        `json:"code"`
	Holiday map[string]*timorYearEntry `json:"holiday"`
}

type timorYearEntry struct {
	Holiday bool   `json:"holiday"`
	Name    string `json:"name"`
	Date    string `json:"date"`
}
