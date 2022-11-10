let ComponentAuth = {
  data() {
    return {
      count: 0
    }
  },
  template: `
    Some text
    <button @click="count++">
      You clicked me {{ count }} times.
    </button>`
}

//#region Main code

Vue.createApp({
  delimiters: ['[[', ']]'],
  components: {
    ComponentAuth: ComponentAuth
  },
  data() {
    return {
      message: 'Hello Vue!'
    }
  }
}).mount('#app')

//#endregion
