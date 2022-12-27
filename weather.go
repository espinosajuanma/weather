package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	API_URL string = "https://api.met.no/weatherapi/locationforecast/2.0"
)

type Request struct {
	params *url.Values
	Agent  string
}

type Response struct {
	Expires      time.Time
	LastModified time.Time
	Body         WeatherResponse
}

func NewRequest() *Request {
	return &Request{
		params: &url.Values{},
	}
}

func (r *Request) SetAgent(a string) { r.Agent = a }

func (r *Request) Get() (Response, error) {
	var res *http.Response

	url := fmt.Sprintf("%s/compact", API_URL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Response{}, err
	}
	req.URL.RawQuery = r.params.Encode()

	if r.Agent != "" {
		req.Header.Set("User-Agent", r.Agent)
	}

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)

	response := Response{}

	var wr WeatherResponse
	if res.Status == "200 OK" {
		response.Expires, _ = time.Parse(time.RFC1123, res.Header.Get("Expires"))
		response.LastModified, _ = time.Parse(time.RFC1123, res.Header.Get("Last-Modified"))

		err := json.Unmarshal(bytes, &wr)
		if err != nil {
			return Response{}, fmt.Errorf("Error trying to unmarshal")
		}
		response.Body = wr
		return response, nil
	} else {
		return Response{}, fmt.Errorf("[%s] status code error", res.Status)
	}

	return Response{}, nil
}

func (r *Request) SetLatitude(lat string) { r.params.Add("lat", lat) }

func (r *Request) SetLongitude(lat string) { r.params.Add("lon", lat) }

func (r *Request) SetCoordinates(lat, lon string) {
	r.SetLatitude(lat)
	r.SetLongitude(lon)
}

func (r *Response) GetCurrent(celsius bool) float64 {
	val := r.Body.Properties.Timeseries[0].Data.Instant.Details.AirTemperature
	if celsius {
		return val
	} else {
		f := (float64(val) * 1.8) + 32
		return f
	}
}

func (r *Response) GetFormat(celsius bool) string {
	sym := "°C"
	if !celsius {
		sym = "°F"
	}
	return fmt.Sprintf("%v %s", r.GetCurrent(celsius), sym)
}
