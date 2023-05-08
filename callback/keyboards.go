package callback

var keyboards = map[string]keyboard{
	"start": {
		Inline:  false,
		OneTime: false,
		Buttons: [][]button{
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "картинка для поднятия настроения",
						Payload: "{\"button\": \"picture\"}",
					},
					Color: "positive",
				},
			},
			{

				{
					Action: action{
						Type:    "callback",
						Label:   "мотивация",
						Payload: "{\"button\": \"motivation\"}",
					},
					Color: "positive",
				},
			},
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "поздравление",
						Payload: "{\"button\": \"greeting\"}",
					},
					Color: "positive",
				},
			},
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "музыка",
						Payload: "{\"button\": \"music\"}",
					},
					Color: "positive",
				},
			},
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "предсказать будущее",
						Payload: "{\"button\": \"future\"}",
					},
					Color: "positive",
				},
			},
		},
	},

	"picture": {
		Inline:  false,
		OneTime: false,
		Buttons: [][]button{
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "котик",
						Payload: "{\"button\": \"cat\"}",
					},
					Color: "positive",
				},
				{
					Action: action{
						Type:    "callback",
						Label:   "собачка",
						Payload: "{\"button\": \"dog\"}",
					},
					Color: "positive",
				},
			},
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "в главное меню",
						Payload: "{\"button\": \"start\"}",
					},
					Color: "primary",
				},
			},
		},
	},

	"motivation": {
		Inline:  false,
		OneTime: false,
		Buttons: [][]button{
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "мотивирующая цитата",
						Payload: "{\"button\": \"quote\"}",
					},
					Color: "positive",
				},
				{
					Action: action{
						Type:    "callback",
						Label:   "пожелать удачи",
						Payload: "{\"button\": \"good luck\"}",
					},
					Color: "positive",
				},
			},
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "в главное меню",
						Payload: "{\"button\": \"start\"}",
					},
					Color: "primary",
				},
			},
		},
	},

	"greeting": {
		Inline:  false,
		OneTime: false,
		Buttons: [][]button{
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "день рождения",
						Payload: "{\"button\": \"birthday\"}",
					},
					Color: "positive",
				},
				{
					Action: action{
						Type:    "callback",
						Label:   "новый год",
						Payload: "{\"button\": \"new year\"}",
					},
					Color: "positive",
				},
			},
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "в главное меню",
						Payload: "{\"button\": \"start\"}",
					},
					Color: "primary",
				},
			},
		},
	},

	"future": {
		Inline:  false,
		OneTime: false,
		Buttons: [][]button{
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "Да, конечно!",
						Payload: "{\"button\": \"thanks\"}",
					},
					Color: "positive",
				},
				{
					Action: action{
						Type:    "callback",
						Label:   "Нет, не знаю",
						Payload: "{\"button\": \"again\"}",
					},
					Color: "positive",
				},
			},
			{
				{
					Action: action{
						Type:  "open_link",
						Link:  "https://github.com/andrew55516",
						Label: "github",
					},
				},
			},
			{
				{
					Action: action{
						Type:    "callback",
						Label:   "в главное меню",
						Payload: "{\"button\": \"start\"}",
					},
					Color: "primary",
				},
			},
		},
	},

	"music": {
		Inline:  true,
		OneTime: false,
		Buttons: [][]button{
			{
				{
					Action: action{
						Type:  "open_link",
						Link:  "https://music.yandex.ru/users/andryusha.axenoff/playlists/1010",
						Label: "спокойная",
					},
				},
			},
			{
				{
					Action: action{
						Type:  "open_link",
						Link:  "https://music.yandex.ru/users/andryusha.axenoff/playlists/1002",
						Label: "заряд энергии",
					},
				},
			},
			{
				{
					Action: action{
						Type:  "open_link",
						Link:  "https://music.yandex.ru/users/andryusha.axenoff/playlists/1005",
						Label: "что-то среднее",
					},
				},
			},
		},
	},
}
