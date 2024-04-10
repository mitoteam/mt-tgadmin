//#region Components

let ComponentMessage = {
  props: ["m", "replymode"],

  template: `
<div class="card">
  <div class="card-header">
    <div class="d-flex justify-content-between">
      <div>
        <span class="fw-bold">{{m.user}}</span><span class="small text-muted ms-3">{{m.date}}</span>
      </div>
      <div>
        <a :class="['btn-' + (replymode ? 'warning' : 'primary')]" class="btn btn-sm" @click="$parent.set_reply_message(replymode ? null : m);" v-text="replymode ? 'Unset' : 'Reply'"></a>
      </div>
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
      reply: null,

      //https://ckeditor.com/ckeditor-5/online-builder/
      editor: ClassicEditor,
		  editorData: '',
		  editorConfig: {},
    }
  },

  components: {
    ComponentMessage: ComponentMessage,
  },

  methods: {
    set_reply_message: function (message) {
      //alert('set_reply_message'); console.log(message);

      this.reply = message;
    },

    logout: function () {
      //alert('logout');
      ApiRequest('logout', null, this, function (response) {
        //console.log(response);
        MtData.session = false; //adjust GUI
      });
    },

    say: function () {
      let data = {
        message: this.editorData,
        reply_to: this.reply?.message_id,
        silent: document.getElementById('silentCheck')?.checked ? 1 : 0,
      }

      //console.log(data); return;

      if(data.message)
      {
        ApiRequest('say', data, this);

        this.editorData = "";
        this.reply = null;

        MtData.status.kind = "info";
        MtData.status.body = "";
      }
      else
      {
        //console.log('Empty set!');
        MtData.status.kind = 'warning';
        MtData.status.body = 'Empty message!';
      }
    },

    list_messages: function () {
      ApiRequest('list_messages', null, this, function (response) {
        //console.log(response);

        if(response.status == "ok")
        {
          //console.log(response.list);
          this.messages = response.list;

          if(response.list.length == 0)
          {
            MtData.status.kind = "info";
            MtData.status.body = "Empty messages list received (" + (new Date()).toLocaleString() + ").";
          }
        }
      });
    },
  },

  template: `
<div id="actions" class="card mb-3">
  <div class="card-body">
    <div class="d-flex justify-content-between">
      <div>
        <a class="btn btn-primary" @click="list_messages();">Get Latest Messages</a>
      </div>
      <div>
        <a class="btn btn-secondary" @click="logout();">Logout</a>
      </div>
    </div>
  </div>
</div>

<div id="send" class="card">
  <div class="card-body">
    <div v-if="reply" class="mb-3">
      <h5 class="card-title">In Reply to:</h5>
      <component-message v-bind:m="reply" v-bind:replymode="true"></component-message>
    </div>

    <h5 class="card-title">Say:</h5>
    <ckeditor :editor="editor" v-model="editorData" :config="editorConfig" id="messageEditor"></ckeditor>
    <div class="form-check">
      <input class="form-check-input" type="checkbox" value="" id="silentCheck">
      <label class="form-check-label" for="silentCheck">
        Silent Message
      </label>
    </div>
    <a class="btn btn-success mt-3" @click="say()">Say</a>
  </div>
</div>

<div id="messages" class="card mt-3" v-if="messages.length > 0">
  <div class="card-body">
    <h5 class="card-title">Latest messages from chat</h5>
    <component-message v-bind:m="m" v-for="m in messages" :key="m.message_id" v-bind:replymode="false"></component-message>
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
<div v-if="body" class="mb-3 alert" :class="['alert-' + kind]" role="alert">
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
.use( CKEditor )
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
      else if(response.status == "error")
      {
        MtData.status.kind = "danger";
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
