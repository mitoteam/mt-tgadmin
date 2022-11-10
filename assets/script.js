//#region Components
let ComponentAuth = {
  data() {
    return {
      count: 0
    }
  },
  template: `
<form>
  <div class="mb-3">
    <label for="password" class="form-label">Password please</label>
    <input type="password" class="form-control" id="password">
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
</form>
`
}
//#endregion

//#region Main code

Vue.createApp({
  delimiters: ['[[', ']]'],
  components: {
    ComponentAuth: ComponentAuth
  },
  data() {
    return {
      session: true
    }
  }
}).mount('#app')

//#endregion
