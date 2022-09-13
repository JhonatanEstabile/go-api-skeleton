package configs

type AuthData struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

var AuthTokens map[string]AuthData = make(map[string]AuthData)

func setTokens() {
	AuthTokens["32c5a206ee876f4c6e1c483457561dbed02a531a89b380c3298bb131a844ac3c"] = AuthData{
		Name:   "app-test",
		Secret: "a1c5930d778e632c6684945ca15bcf3c752d17502d4cfbd1184024be6de14540",
	}

	AuthTokens["00e1479a7d2f87d9ef4aa400e0a64a64959b4cd7ec0bee9a08df1ba0c47f9ad7"] = AuthData{
		Name:   "theio-app",
		Secret: "dc57d74987de604218f4f5a57f7115a5832ad8a26ae89a77420994d01dc9945a",
	}
}

func GetTokens() map[string]AuthData {
	if len(AuthTokens) == 0 {
		setTokens()
	}

	return AuthTokens
}
