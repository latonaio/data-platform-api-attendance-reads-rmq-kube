package requests

type Header struct {
	Attendance				int		`json:"Attendance"`
	AttendanceDate			string	`json:"AttendanceDate"`
	AttendanceTime			string	`json:"AttendanceTime"`
	Attender				int		`json:"Attender"`
	AttendanceObjectType	string	`json:"AttendanceObjectType"`
	AttendanceObject		int		`json:"AttendanceObject"`
	Participation			*int	`json:"Participation"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	IsCancelled				*bool	`json:"IsCancelled"`
}
