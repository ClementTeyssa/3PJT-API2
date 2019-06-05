package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

const ApiKey = "KlM5ybPPkwzKoYmK7zNv8VTlFeeBoi+bts8m7Cfoa6V60ThC1UtFAms41ZPg+NxPqI33+t96cPs6zL2OiuxcZ0wWTcRMu/dxYTKMiQB24qNvImcusUQ7KazN9jbNsBaqo44OKF0SLMpkVB5HbBi3bgGCvtesZONYOl1Bjwq7GgZzb+PEfnHC/D7hDQpGXvvSRPi45HREz24+8NoH8ISIK9VUKY3MZPXoA5FcWopQPTlNXUkF4QKxPjHept5J8xBOQoxw4hHSWvDfLtAoGBAP+6CtPeFmmHLBWTtSGHG8nZAjnQJSBjGFqbkq9elRvC8to52etDOeqoypqbl84Z0t2lDq3l1V1g1L+EOp5oCmi30o1znrUxn+36oS32KM/Z8paHMY4odiKu5Qd+dXjtkanEizsheCB3Qo6SOIeYz+C81ycmZ59onwY13fiSDtBjAoGBAOGYzGGn5rycjAiY2JMqoBGgcnlEmeH5l2UBbKrZinhoPwyTc1Tl0lj3PFo9VdA3pzYo1EpvcSXNSBEk+3iFtd08YCWQlM+BOfISB2kpcjGJpRE0cBk+H6P+GYrkzVAnbk5XDHeQpHbtHFFe5W7y7mFGQqF68tu52ljn2lYeknCLAoGBAPDZK7F3jew93wtDXmkBcu6ccDg7DXx/WESeGX0Ju6214jweGFw3qKiV6JvMxyDZuXv/JOArgtS7iiQGStt2UMnYs+v3FWZSGvbUshAoGBAPFGUCgRLOOLBf5c2421kWwQw1pOZt3sfmjjKbwM9huJVa6R7evdGL5ZfGRPWIjqcd1q/ZtJi3Ocjt9zxp7YHx+8WKIFzYicmy2V3/8PFrNPvQc9JD0dnqAZWKPjbdeP/hX+/YM03ckd15vYwhVGOhh5Bfp7XJXOeXHA5AEPMiZbAoGBAPhtnFNh6Zt5uWmUzifZsWKh7OlTf1zfRbtP5bTLaG1rLYhiqB+LRYQxJkY0fGYFggw6dPZQNYkxss7skPjVU23ImJdd6KgK/oA9QbaGetDaObShxOT9QiNiulpvBE3mfTehEcb8fy3DO7XxwVu///QWQa/fcRVu1tfE6r40ypnZAoGBAJbWqo4iDrpC1ABt9STm+n9hC4CNu2nKzEyumAzk3YfhmBuHSANJOormTwo3QNZ4G6+dj+dhVzl0Cy5Pp0DvRKYGjjLdBd3+qkPUOOuXOyYjCVwMLeCwECgYEA5KUH1kX4t4F55xchXEw7EGm5n5Yb+/3fbMq1r8esdbGJW5PVNdU1S7DvPVk6OW1QhFRUuWUWruNcXZx+gh9j7lL+YJ395CiuPWb1pSmEagkMh+4y6LCcnA8N5jlyGz+BsBBdqB4YzvHe6EN0MRhz/p4BjOxTY/ozxkGQgpERqr8CgYEAw7VCU5WvpYLbWDRVoXeRQQsDqATiV4u3m6s22I7addx/1okp4wxMuF3lzQ1Ik4K0dPXrU+qMvSVyHkIKtGjzb9lynD4oUn9cd6CbbmYYrWJq108BzWQC5H8LLGQjWykLqPnFuzAXUAboFzFxt0oaiIgRrTlGGi3rcPa1jrb59pUCgYAOROZaTv+YpDHxZBYm5nMBMTfuZ7XTYiKhbX0tPS7cIXpLBszVkTrEQISNtZXWo6XUvFaVQ2PJAlNvXCbhBAh+hRtu+C0yIf/NwTEvX4"

type MyError struct {
	Error string `json:"error"`
}

func LogRequest(r *http.Request) {
	log.Println("Request to " + r.URL.String() + " with " + r.Method + " method")
}

func ErrorHandlerHttpRespond(w http.ResponseWriter, err string) {
	log.Println("\n---------------ERROR---------------\n" + err + "\n---------------ERROR---------------")
	var error MyError
	error.Error = err
	json.NewEncoder(w).Encode(error)
}
