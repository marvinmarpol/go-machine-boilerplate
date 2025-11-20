package lockmanager

import (
	"time"
)

type LockManager struct {
	MapLock map[string]Resource
}

type Resource struct {
	ClientID string
	Expired  time.Time
}

func NewLockManager() *LockManager {
	return &LockManager{
		MapLock: make(map[string]Resource),
	}
}

func (l *LockManager) Lock(resourceId, clientId string, ttl time.Duration) bool {
	currentTime := time.Now()
	resource, ok := l.MapLock[resourceId]

	if ok && currentTime.Before(resource.Expired) {
		return false
	}

	l.MapLock[resourceId] = Resource{
		ClientID: clientId,
		Expired:  time.Now().Add(ttl),
	}
	return true
}

func (l *LockManager) Unlock(resourceId, clientId string) bool {
	currentTime := time.Now()
	resource, ok := l.MapLock[resourceId]

	if ok && resource.ClientID == clientId {
		delete(l.MapLock, resourceId)

		if currentTime.Before(resource.Expired) {
			return true
		}
	}

	return false
}
