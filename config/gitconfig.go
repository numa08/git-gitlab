package config


type GitConfig interface  {
    Host() (string, error)
    ApiPath() (string, error)
    Token() (string, error)
}