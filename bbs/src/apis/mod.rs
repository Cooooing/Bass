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

pub mod account_service_api;
pub mod article_service_api;
pub mod auth_service_api;
pub mod comment_service_api;
pub mod domain_service_api;
pub mod location_service_api;
pub mod notification_service_api;
pub mod postscript_service_api;
pub mod preferences_service_api;
pub mod privacy_setting_service_api;
pub mod relation_service_api;
pub mod tag_service_api;
pub mod tfa_service_api;

pub mod configuration;

use std::sync::Arc;

pub trait Api {
    fn account_service_api(&self) -> &dyn account_service_api::AccountServiceApi;
    fn article_service_api(&self) -> &dyn article_service_api::ArticleServiceApi;
    fn auth_service_api(&self) -> &dyn auth_service_api::AuthServiceApi;
    fn comment_service_api(&self) -> &dyn comment_service_api::CommentServiceApi;
    fn domain_service_api(&self) -> &dyn domain_service_api::DomainServiceApi;
    fn location_service_api(&self) -> &dyn location_service_api::LocationServiceApi;
    fn notification_service_api(&self) -> &dyn notification_service_api::NotificationServiceApi;
    fn postscript_service_api(&self) -> &dyn postscript_service_api::PostscriptServiceApi;
    fn preferences_service_api(&self) -> &dyn preferences_service_api::PreferencesServiceApi;
    fn privacy_setting_service_api(&self) -> &dyn privacy_setting_service_api::PrivacySettingServiceApi;
    fn relation_service_api(&self) -> &dyn relation_service_api::RelationServiceApi;
    fn tag_service_api(&self) -> &dyn tag_service_api::TagServiceApi;
    fn tfa_service_api(&self) -> &dyn tfa_service_api::TfaServiceApi;
}

pub struct ApiClient {
    account_service_api: Box<dyn account_service_api::AccountServiceApi>,
    article_service_api: Box<dyn article_service_api::ArticleServiceApi>,
    auth_service_api: Box<dyn auth_service_api::AuthServiceApi>,
    comment_service_api: Box<dyn comment_service_api::CommentServiceApi>,
    domain_service_api: Box<dyn domain_service_api::DomainServiceApi>,
    location_service_api: Box<dyn location_service_api::LocationServiceApi>,
    notification_service_api: Box<dyn notification_service_api::NotificationServiceApi>,
    postscript_service_api: Box<dyn postscript_service_api::PostscriptServiceApi>,
    preferences_service_api: Box<dyn preferences_service_api::PreferencesServiceApi>,
    privacy_setting_service_api: Box<dyn privacy_setting_service_api::PrivacySettingServiceApi>,
    relation_service_api: Box<dyn relation_service_api::RelationServiceApi>,
    tag_service_api: Box<dyn tag_service_api::TagServiceApi>,
    tfa_service_api: Box<dyn tfa_service_api::TfaServiceApi>,
}

impl ApiClient {
    pub fn new(configuration: Arc<configuration::Configuration>) -> Self {
        Self {
            account_service_api: Box::new(account_service_api::AccountServiceApiClient::new(configuration.clone())),
            article_service_api: Box::new(article_service_api::ArticleServiceApiClient::new(configuration.clone())),
            auth_service_api: Box::new(auth_service_api::AuthServiceApiClient::new(configuration.clone())),
            comment_service_api: Box::new(comment_service_api::CommentServiceApiClient::new(configuration.clone())),
            domain_service_api: Box::new(domain_service_api::DomainServiceApiClient::new(configuration.clone())),
            location_service_api: Box::new(location_service_api::LocationServiceApiClient::new(configuration.clone())),
            notification_service_api: Box::new(notification_service_api::NotificationServiceApiClient::new(configuration.clone())),
            postscript_service_api: Box::new(postscript_service_api::PostscriptServiceApiClient::new(configuration.clone())),
            preferences_service_api: Box::new(preferences_service_api::PreferencesServiceApiClient::new(configuration.clone())),
            privacy_setting_service_api: Box::new(privacy_setting_service_api::PrivacySettingServiceApiClient::new(configuration.clone())),
            relation_service_api: Box::new(relation_service_api::RelationServiceApiClient::new(configuration.clone())),
            tag_service_api: Box::new(tag_service_api::TagServiceApiClient::new(configuration.clone())),
            tfa_service_api: Box::new(tfa_service_api::TfaServiceApiClient::new(configuration.clone())),
        }
    }
}

impl Api for ApiClient {
    fn account_service_api(&self) -> &dyn account_service_api::AccountServiceApi {
        self.account_service_api.as_ref()
    }
    fn article_service_api(&self) -> &dyn article_service_api::ArticleServiceApi {
        self.article_service_api.as_ref()
    }
    fn auth_service_api(&self) -> &dyn auth_service_api::AuthServiceApi {
        self.auth_service_api.as_ref()
    }
    fn comment_service_api(&self) -> &dyn comment_service_api::CommentServiceApi {
        self.comment_service_api.as_ref()
    }
    fn domain_service_api(&self) -> &dyn domain_service_api::DomainServiceApi {
        self.domain_service_api.as_ref()
    }
    fn location_service_api(&self) -> &dyn location_service_api::LocationServiceApi {
        self.location_service_api.as_ref()
    }
    fn notification_service_api(&self) -> &dyn notification_service_api::NotificationServiceApi {
        self.notification_service_api.as_ref()
    }
    fn postscript_service_api(&self) -> &dyn postscript_service_api::PostscriptServiceApi {
        self.postscript_service_api.as_ref()
    }
    fn preferences_service_api(&self) -> &dyn preferences_service_api::PreferencesServiceApi {
        self.preferences_service_api.as_ref()
    }
    fn privacy_setting_service_api(&self) -> &dyn privacy_setting_service_api::PrivacySettingServiceApi {
        self.privacy_setting_service_api.as_ref()
    }
    fn relation_service_api(&self) -> &dyn relation_service_api::RelationServiceApi {
        self.relation_service_api.as_ref()
    }
    fn tag_service_api(&self) -> &dyn tag_service_api::TagServiceApi {
        self.tag_service_api.as_ref()
    }
    fn tfa_service_api(&self) -> &dyn tfa_service_api::TfaServiceApi {
        self.tfa_service_api.as_ref()
    }
}


