package output

type DayCount struct {
	Day   string `json:"day" db:"day"`
	Count int64  `json:"count" db:"count"`
}

type DeviceCount struct {
	Device string `json:"device" db:"device"`
	Count  int64  `json:"count" db:"count"`
}

type DashboardStats struct {
	TotalProjects    int64         `json:"total_projects"`
	TotalUsers       int64         `json:"total_users"`
	ActiveSessions   int64         `json:"active_sessions"`
	TotalClients     int64         `json:"total_clients"`
	LoginsPerDay     []DayCount    `json:"logins_per_day"`
	SessionsByDevice []DeviceCount `json:"sessions_by_device"`
}
