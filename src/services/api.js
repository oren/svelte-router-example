const api = {
  login() {
    return {name: 'josh', email: 'josh@gmail.com', token: 'abcde'}
  },
  getDoctors() {
    return new Promise(function(resolve, reject) {
      fetch('https://api.github.com/users/mralexgray/repos', {
        method: 'get'
      }).then(function(response) {

        response.json().then(function(data) {
          console.log('success fetching', data);
          resolve(data);
        });

      }).catch(function(err) {
        console.log('error fetching', err);
        reject(err);
      });
    })

  }
};

export default api;
