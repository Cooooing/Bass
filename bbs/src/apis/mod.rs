use std::error;
use std::fmt;

#[derive(Debug, Clone)]
pub struct ResponseContent<T> {
    pub status: reqwest::StatusCode,
    pub content: String,
    pub entity: Option<T>,
}

#[derive(Debug)]
pub enum Error<T> {
    Reqwest(reqwest::Error),
    Serde(serde_json::Error),
    Io(std::io::Error),
    ResponseError(ResponseContent<T>),
}

impl <T> fmt::Display for Error<T> {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let (module, e) = match self {
            Error::Reqwest(e) => ("reqwest", e.to_string()),
            Error::Serde(e) => ("serde", e.to_string()),
            Error::Io(e) => ("IO", e.to_string()),
            Error::ResponseError(e) => ("response", format!("status code {}", e.status)),
        };
        write!(f, "error in {}: {}", module, e)
    }
}

impl <T: fmt::Debug> error::Error for Error<T> {
    fn source(&self) -> Option<&(dyn error::Error + 'static)> {
        Some(match self {
            Error::Reqwest(e) => e,
            Error::Serde(e) => e,
            Error::Io(e) => e,
            Error::ResponseError(_) => return None,
        })
    }
}

impl <T> From<reqwest::Error> for Error<T> {
    fn from(e: reqwest::Error) -> Self {
        Error::Reqwest(e)
    }
}

impl <T> From<serde_json::Error> for Error<T> {
    fn from(e: serde_json::Error) -> Self {
        Error::Serde(e)
    }
}

impl <T> From<std::io::Error> for Error<T> {
    fn from(e: std::io::Error) -> Self {
        Error::Io(e)
    }
}

pub fn urlencode<T: AsRef<str>>(s: T) -> String {
    ::url::form_urlencoded::byte_serialize(s.as_ref().as_bytes()).collect()
}

pub fn parse_deep_object(prefix: &str, value: &serde_json::Value) -> Vec<(String, String)> {
    if let serde_json::Value::Object(object) = value {
        let mut params = vec![];

        for (key, value) in object {
            match value {
                serde_json::Value::Object(_) => params.append(&mut parse_deep_object(
                    &format!("{}[{}]", prefix, key),
                    value,
                )),
                serde_json::Value::Array(array) => {
                    for (i, value) in array.iter().enumerate() {
                        params.append(&mut parse_deep_object(
                            &format!("{}[{}][{}]", prefix, key, i),
                            value,
                        ));
                    }
                },
                serde_json::Value::String(s) => params.push((format!("{}[{}]", prefix, key), s.clone())),
                _ => params.push((format!("{}[{}]", prefix, key), value.to_string())),
            }
        }

        return params;
    }

    unimplemented!("Only objects are supported with style=deepObject")
}

/// Internal use only
/// A content type supported by this client.
#[allow(dead_code)]
enum ContentType {
    Json,
    Text,
    Unsupported(String)
}

impl From<&str> for ContentType {
    fn from(content_type: &str) -> Self {
        if content_type.starts_with("application") && content_type.contains("json") {
            return Self::Json;
        } else if content_type.starts_with("text/plain") {
            return Self::Text;
        } else {
            return Self::Unsupported(content_type.to_string());
        }
    }
}

pub mod account_service;
pub mod article_service;
pub mod auth_service;
pub mod checkin_service;
pub mod comment_service;
pub mod domain_service;
pub mod location_service;
pub mod notification_service;
pub mod otp_service;
pub mod postscript_service;
pub mod preferences_service;
pub mod privacy_setting_service;
pub mod relation_service;
pub mod tag_service;

pub mod configuration;

use std::sync::Arc;

pub trait Api {
    fn account_service(&self) -> &dyn account_service::AccountService;
    fn article_service(&self) -> &dyn article_service::ArticleService;
    fn auth_service(&self) -> &dyn auth_service::AuthService;
    fn checkin_service(&self) -> &dyn checkin_service::CheckinService;
    fn comment_service(&self) -> &dyn comment_service::CommentService;
    fn domain_service(&self) -> &dyn domain_service::DomainService;
    fn location_service(&self) -> &dyn location_service::LocationService;
    fn notification_service(&self) -> &dyn notification_service::NotificationService;
    fn otp_service(&self) -> &dyn otp_service::OtpService;
    fn postscript_service(&self) -> &dyn postscript_service::PostscriptService;
    fn preferences_service(&self) -> &dyn preferences_service::PreferencesService;
    fn privacy_setting_service(&self) -> &dyn privacy_setting_service::PrivacySettingService;
    fn relation_service(&self) -> &dyn relation_service::RelationService;
    fn tag_service(&self) -> &dyn tag_service::TagService;
}

pub struct ApiClient {
    account_service: Box<dyn account_service::AccountService>,
    article_service: Box<dyn article_service::ArticleService>,
    auth_service: Box<dyn auth_service::AuthService>,
    checkin_service: Box<dyn checkin_service::CheckinService>,
    comment_service: Box<dyn comment_service::CommentService>,
    domain_service: Box<dyn domain_service::DomainService>,
    location_service: Box<dyn location_service::LocationService>,
    notification_service: Box<dyn notification_service::NotificationService>,
    otp_service: Box<dyn otp_service::OtpService>,
    postscript_service: Box<dyn postscript_service::PostscriptService>,
    preferences_service: Box<dyn preferences_service::PreferencesService>,
    privacy_setting_service: Box<dyn privacy_setting_service::PrivacySettingService>,
    relation_service: Box<dyn relation_service::RelationService>,
    tag_service: Box<dyn tag_service::TagService>,
}

impl ApiClient {
    pub fn new(configuration: Arc<configuration::Configuration>) -> Self {
        Self {
            account_service: Box::new(account_service::AccountServiceClient::new(configuration.clone())),
            article_service: Box::new(article_service::ArticleServiceClient::new(configuration.clone())),
            auth_service: Box::new(auth_service::AuthServiceClient::new(configuration.clone())),
            checkin_service: Box::new(checkin_service::CheckinServiceClient::new(configuration.clone())),
            comment_service: Box::new(comment_service::CommentServiceClient::new(configuration.clone())),
            domain_service: Box::new(domain_service::DomainServiceClient::new(configuration.clone())),
            location_service: Box::new(location_service::LocationServiceClient::new(configuration.clone())),
            notification_service: Box::new(notification_service::NotificationServiceClient::new(configuration.clone())),
            otp_service: Box::new(otp_service::OtpServiceClient::new(configuration.clone())),
            postscript_service: Box::new(postscript_service::PostscriptServiceClient::new(configuration.clone())),
            preferences_service: Box::new(preferences_service::PreferencesServiceClient::new(configuration.clone())),
            privacy_setting_service: Box::new(privacy_setting_service::PrivacySettingServiceClient::new(configuration.clone())),
            relation_service: Box::new(relation_service::RelationServiceClient::new(configuration.clone())),
            tag_service: Box::new(tag_service::TagServiceClient::new(configuration.clone())),
        }
    }
}

impl Api for ApiClient {
    fn account_service(&self) -> &dyn account_service::AccountService {
        self.account_service.as_ref()
    }
    fn article_service(&self) -> &dyn article_service::ArticleService {
        self.article_service.as_ref()
    }
    fn auth_service(&self) -> &dyn auth_service::AuthService {
        self.auth_service.as_ref()
    }
    fn checkin_service(&self) -> &dyn checkin_service::CheckinService {
        self.checkin_service.as_ref()
    }
    fn comment_service(&self) -> &dyn comment_service::CommentService {
        self.comment_service.as_ref()
    }
    fn domain_service(&self) -> &dyn domain_service::DomainService {
        self.domain_service.as_ref()
    }
    fn location_service(&self) -> &dyn location_service::LocationService {
        self.location_service.as_ref()
    }
    fn notification_service(&self) -> &dyn notification_service::NotificationService {
        self.notification_service.as_ref()
    }
    fn otp_service(&self) -> &dyn otp_service::OtpService {
        self.otp_service.as_ref()
    }
    fn postscript_service(&self) -> &dyn postscript_service::PostscriptService {
        self.postscript_service.as_ref()
    }
    fn preferences_service(&self) -> &dyn preferences_service::PreferencesService {
        self.preferences_service.as_ref()
    }
    fn privacy_setting_service(&self) -> &dyn privacy_setting_service::PrivacySettingService {
        self.privacy_setting_service.as_ref()
    }
    fn relation_service(&self) -> &dyn relation_service::RelationService {
        self.relation_service.as_ref()
    }
    fn tag_service(&self) -> &dyn tag_service::TagService {
        self.tag_service.as_ref()
    }
}


