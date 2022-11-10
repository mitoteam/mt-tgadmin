//#region Components

let ComponentMain = {
  data() {
    return {
      count: 0
    }
  },
  template: `
  <div class="mb-3">
    <button class="btn btn-primary" @click="$parent.logout();">Logout</button>
  </div>
`
}

let ComponentAuth = {
  data() {
    return {
      count: 0
    }
  },
  methods: {
    password: function () {
      let data = {
        password: document.getElementById('password').value,
      }

      ApiRequest('password', data, this, function (response) {
        //console.log(response);

        if(response.status == "OK")
        {
          this.$parent.session = true; //adjust GUI
        }
        else
        {
          alert(response.message);
          document.getElementById('password').value = '';
        }
      });
    }
  },
  template: `
<div class="mb-3">
  <label for="password" class="form-label">Password please</label>
  <input type="password" class="form-control" id="password">
</div>
<a class="btn btn-success" @click="password();">Authorize</a>
`
}
//#endregion

//#region Main code

Vue.createApp({
  delimiters: ['[[', ']]'],
  components: {
    ComponentMain: ComponentMain,
    ComponentAuth: ComponentAuth
  },
  methods: {
    logout: function () {
      //alert('logout');
      ApiRequest('logout', null, this, function (response) {
        console.log(response);
        this.session = false; //adjust GUI
      });
    }
  },
  data() {
    return {
      session: mtAuth // take initial value from global variable provided in index.html
    }
  }
}).mount('#app')

//#endregion

//#region Helpers
function ApiRequest(path, data, component, responseHandler)
{
  fetch(
    '/api/' + path,
    {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }
  )
    .then(response => response.json())
    .then(response => responseHandler.call(component, response));
}
//#endregion
