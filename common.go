package main


type Reporter interface{
  ReportHealth(h *Health)
}

type ReporterCredentials struct {
  User string `json:"user"`
  Key string `json:"key"`
}

type GroundControlConfig struct {
  Temperature string `json:"temperature"`
  Port int `json:"port"`
  Host string `json:"host"`
  Stdout bool `json:"stdout"`
  Interval int `json:"interval"`
  HistoryInterval int `json:"history_interval"`
  HistoryBacklog int `json:"history_backlog"`
  Librato ReporterCredentials `json:"librato"`
  TempoDB ReporterCredentials `json:"tempodb"`
  Controls map[string]interface{} `json:"controls"`
}

