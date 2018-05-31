package operation

import (
	"errors"
	"strconv"
	"time"
)

type VidInfo struct {
	Locations       []Location
	NextRefreshTime time.Time
}
type VidCache struct {
	cache []VidInfo
}

func (vc *VidCache) Get(vid string) ([]Location, error) {
	id, _ := strconv.Atoi(vid)
	if 0 < id && id <= len(vc.cache) {
		if vc.cache[id-1].Locations == nil {
			return nil, errors.New("Not Set")
		}
		if vc.cache[id-1].NextRefreshTime.Before(time.Now()) {
			return nil, errors.New("Expired")
		}
		return vc.cache[id-1].Locations, nil
	}
	return nil, errors.New("Not Found")
}
func (vc *VidCache) Set(vid string, locations []Location, duration time.Duration) {
	id, _ := strconv.Atoi(vid)
	if id >= len(vc.cache) {
		for i := id - len(vc.cache); i > 0; i-- {
			vc.cache = append(vc.cache, VidInfo{})
		}
	}

	vc.cache[id-1].Locations = locations
	vc.cache[id-1].NextRefreshTime = time.Now().Add(duration)

}
