<template>
    <div id = "WelcomeDiv">
      <div class="container-fluid small-fluid-container">
        <form id="WelcomeForm" class="small-center-form" v-on:submit.prevent>
          <h1 class="card-header">Rechnungscockpit</h1>
          <hr/>
          <div class="input-group mt-1">
              <span class="input-group-text" id="basic-addon1">username</span>
              <input type="text" class="form-control" v-tooltip="{ value: 'Enter a Username', showDelay: 1000, hideDelay: 300}" aria-label="Username" v-model="username">
          </div>
          <div class>
            <button v-if="this.validateInput() == true" v-on:click="dummy1" v-tooltip="{ value: 'Create a new Game', showDelay: 1000, hideDelay: 300 }" class="btn btn-primary-lavender w-75" type="Submit" id="new-game" aria-expanded="false">New Game</button>
            <button v-else v-on:click="dummy1" v-tooltip="{ value: 'Please enter username',pt: {text: 'text-danger'}}" class="btn btn-primary-lavender w-75" type="Submit" id="new-game" aria-expanded="false">New Game</button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  //import SocketioService from '../../services/socketio.service.js';
  
  export default {
    name: 'LandingPage',
  
    created(){
      /*if(SocketioService.socket.connected == false){
        console.log("Hello there!");
        SocketioService.connect();
      }*/
    },
  
    data() {
      return {
        username: null,
        intent: null,
        game: null,
        popover1: "display: none",
        popover2: "display: none",
        popover3: "display: none",
      };
    },
  
    methods: {
      validateInput() {
        const regex = /^[a-zA-Z0-9]+$/;
  
        if (!this.username) {
          //add error handling -> user input needed
          return;
        }
  
        if (!regex.test(this.username)) {
          return;
        }
  
        return true;
      }
  
    }
  }
  </script>
  
  <style scoped>
   
  </style>