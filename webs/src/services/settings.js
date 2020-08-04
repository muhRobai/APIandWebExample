import Axios from 'axios';

class Settings {
	config() {
		return {
				timeout: 1000,
		}
	}

	post(path, payload) {
		return Axios.post(path, payload);
	}

	get(path) {
		return Axios.get(path, this.config());
	}

	getData() {
		return new Promise((resolve, reject) => {
			this.get('https://jsonplaceholder.typicode.com/todos').then((v) => {
				resolve(v);
			}).catch((err) => {
				console.log(err);
				reject(err);
			})
		})
	}
}

export default Settings;