const api = {
  login() {
    return {name: 'josh', email: 'josh@gmail.com'}
  },
  getDoctors() {
    return [
      {name: 'Wong'},
      {name: 'Tan'},
      {name: 'Ruth'}
    ]
  }
};

export default api;
