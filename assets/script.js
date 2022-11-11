//#region Components

let ComponentMain = {
  data() {
    return {
      message: ""
    }
  },

  methods: {
    logout: function () {
      //alert('logout');
      ApiRequest('logout', null, this, function (response) {
        //console.log(response);
        MtData.session = false; //adjust GUI
      });
    },

    say: function () {
      let data = {
        message: document.getElementById('messageEditor').value,
      }

      //console.log(data); return;

      ApiRequest('say', data, this, function (response) {
        //console.log(response);

        if(response.status == "ok")
        {

        }
      });
    },
  },

  template: `
<div id="actions" class="card mb-3">
  <div class="card-body">
    <div class="d-flex justify-content-between">
      <div>
        no actions yet
      </div>
      <div>
        <a class="btn btn-secondary" @click="logout();">Logout</a>
      </div>
    </div>
  </div>
</div>

<div>
  <label for="messageEditor" class="form-label fw-bold">Message:</label>
  <textarea class="form-control" id="messageEditor" rows="10" placeholder="Message text"></textarea>
  <a class="btn btn-success mt-3" @click="say()">Say</a>
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
  <input type="password" class="form-control" id="password" @keyup.enter="password();">
</div>
<a class="btn btn-success" @click="password();">Authorize</a>
`
}
//#endregion

//#region Main code

var MtData = {
  // take initial value from global variable provided in index.html
  session: mtAuth,
  message: { body: "", kind: "primary" },
}

Vue.createApp({
  // delimiters are set only for this component (index.html is go template). Each component has it own delimiters.
  delimiters: ['[[', ']]'],

  components: {
    ComponentMessage: ComponentMessage,
    ComponentMain: ComponentMain,
    ComponentAuth: ComponentAuth
  },

  data() {
    MtData = Vue.reactive(MtData);
    return MtData;
  }
})
.mount('#app')

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
