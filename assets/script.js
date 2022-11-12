//#region Components

let ComponentMessage = {
  props: ["m"],

  template: `
<div class="card">
  <div class="card-header">
    <div class="d-flex justify-content-between">
      <div class="fw-bold">{{m.user}}</div>
      <div>{{m.date}}</div>
    </div>
  </div>
  <div class="card-body">
    {{m.message}}
  </div>
</div>`
}

let ComponentMain = {
  data() {
    return {
      messages: [],
    }
  },

  components: {
    ComponentMessage: ComponentMessage,
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

      ApiRequest('say', data, this);
    },

    list_messages: function () {
      ApiRequest('list_messages', null, this, function (response) {
        //console.log(response);

        if(response.status == "ok")
        {
          console.log(response.list);
          this.messages = response.list;
        }
      });
    },
  },

  template: `
<div id="actions" class="card mb-3">
  <div class="card-body">
    <div class="d-flex justify-content-between">
      <div>
        <a class="btn btn-primary" @click="list_messages();">Update Messages</a>
      </div>
      <div>
        <a class="btn btn-secondary" @click="logout();">Logout</a>
      </div>
    </div>
  </div>
</div>

<div id="messages" class="card mb-3" v-if="messages.length > 0">
  <div class="card-body">
    <h5 class="card-title">Messages from chat</h5>
    <component-message v-bind:m="m" v-for="m in messages" :key="m.message_id"></component-message>
  </div>
</div>

<div id="messages" class="card">
  <div class="card-body">
    <h5 class="card-title">Say:</h5>
    <textarea class="form-control" id="messageEditor" rows="10" placeholder="Message text"></textarea>
    <a class="btn btn-success mt-3" @click="say()">Say</a>
  </div>
</div>
`
}

let ComponentStatus = {
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
  status: { body: "", kind: "primary" },
}

Vue.createApp({
  // delimiters are set only for this component (index.html is go template). Each component has it own delimiters.
  delimiters: ['[[', ']]'],

  components: {
    ComponentStatus: ComponentStatus,
    ComponentAuth: ComponentAuth,
    ComponentMain: ComponentMain
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
        MtData.status.kind = "success";
      }
      else
      {
        MtData.status.kind = response.status;
      }

      MtData.status.body = response.message ?? "";

      if(typeof(responseHandler) == 'function')
      {
        responseHandler.call(component, response);
      }
    });
}
//#endregion
