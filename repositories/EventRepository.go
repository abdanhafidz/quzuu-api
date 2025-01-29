package repositories

import (
	"github.com/quzuu-be/models"
)

type EventRepositoryWrapper struct {
	SingleWrapper models.Events
	ArrayWrapper  []models.Events
}

var EventRepository = Repository[models.Events]{}
var EventPaginateRepository = EventRepositoryWrapper{
	SingleWrapper: models.Events{},
	ArrayWrapper:  []models.Events{},
}
var EventAssignRepository = Repository[models.EventAssign]{}

func (event *EventRepositoryWrapper) FindAllPaginate(offset int, limit int, filter string, id_account int) RepositoryResult[[]models.Events] {
	rows := db.Raw("(SELECT events.id_event,events.title, events.start_event, events.end_event, events.s_id, events.public FROM events WHERE public = 'Y' ) UNION (SELECT events.id_event,events.title, events.start_event, events.end_event, events.s_id, events.public FROM events INNER JOIN event_assign ON events.id_event=event_assign.id_event WHERE event_assign.id_account = ? )", id_account).Limit(limit).Offset(offset).Where("title LIKE ? OR 1=1", filter).Find(&event.ArrayWrapper)
	return RepositoryResult[[]models.Events]{
		Result:    event.ArrayWrapper,
		RowsCount: int(rows.RowsAffected),
		NoRecord:  bool(int(rows.RowsAffected) == 0),
		RowsError: rows.Error,
	}
}
