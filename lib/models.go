package lib

import (
)

type Project struct{

    project_id  int8
    name        string
    description string
}

type Service struct{
    service_id  int8
    name        string
    lang        string
    focus       string
    project     int
}

