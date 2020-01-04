package notification


import(
	
	"github.com/robi_a21/Cassiopeia/entity"
)
type NotificationRepository interface{

	 GetNotification() ([]entity.Notification ,error)
     AddNotification( entity.Notification)error

}