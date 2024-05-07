package dpfm_api_output_formatter

import (
	"data-platform-api-attendance-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Attendance,
			&pm.AttendanceDate,
			&pm.AttendanceTime,
			&pm.Attender,
			&pm.AttendanceObjectType,
			&pm.AttendanceObject,
			&pm.Participation,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Attendance:				data.Attendance,
			AttendanceDate:			data.AttendanceDate,
			AttendanceTime:			data.AttendanceTime,
			Attender:				data.Attender,
			AttendanceObjectType:	data.AttendanceObjectType,
			AttendanceObject:		data.AttendanceObject,
			Participation:			data.Participation,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			IsCancelled:			data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
