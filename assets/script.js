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

let ComponentMessage = {
  props: {
    body: String,
    kind: {
      type: String,
      default: "primary",
    },
  },

  template: `
<div v-if="body" :class="['mb-3', 'alert', 'alert-' + kind]" role="alert">
  {{body}}
</div>`
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

        if(response.status == "ok")
        {
          MtData.session = true; //adjust GUI
        }
        else
        {
          //alert(response.message);
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

var MtData = {
  // take initial value from global variable provided in index.html
  session: mtAuth,
  message: { body: "", kind: "primary" }
}

Vue.createApp({
  // delimiters are set only for this component (index.html is go template). Each component has it own delimiters.
  delimiters: ['[[', ']]'],
  components: {
    ComponentMessage: ComponentMessage,
    ComponentMain: ComponentMain,
    ComponentAuth: ComponentAuth
  },
  methods: {
    logout: function () {
      //alert('logout');
      ApiRequest('logout', null, this, function (response) {
        //console.log(response);
        MtData.session = false; //adjust GUI
      });
    }
  },
  data() {
    MtData = Vue.reactive(MtData);
    return MtData;
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
      body: JSON.stringify(data ?? {})
    }
  )
    .then(response => response.json())
    .then(function(response){
      if(response.status == "ok")
      {
        MtData.message.kind = "success";
      }
      else
      {
        MtData.message.kind = response.status;
      }

      MtData.message.body = response.message ?? "";

      responseHandler.call(component, response);
    });
}
//#endregion
