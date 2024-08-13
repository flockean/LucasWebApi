package lib

import (
)

type Project struct{

    project_id  int8
    name        string
    description string
    services    []Service
}

type Service struct{
    service_id  int8
    name        string
    lang        string
    focus       string
}

