const api = {
  login() {
    return {name: 'josh', email: 'josh@gmail.com', token: 'abcde'}
  },
  getDoctors() {
    fetch('https://api.github.com/users/mralexgray/repos', {
      method: 'get'
    }).then(function(response) {
      console.log('success fetching', response);
      return [
        {name: 'Wong'},
        {name: 'Tan'},
        {name: 'Ruth'}
      ];
    }).catch(function(err) {
      console.log('error fetching', err);
      return [];
    });

  }
};

export default api;
