<template>
  <div>
    <div v-if="response">
      <h2>{{ rHeader }}:</h2>
      <p class="yell">{{ response }}</p>
      <div v-if="rHeader === 'You yelled'">
        <h3>
          at:
        </h3>
        <p class="yell">{{ pCloud }}</p>
      </div>
      <button v-on:click="again">Yell Again</button>
    </div>
    <div v-else>
      <textarea name="msg" id="msg" v-model="msg">
      </textarea>
      <select name="cloud" id="cloud" v-model="cloud">
        <option value="">--Please Choose a Cloud</option>
        <option value="google">Google</option>
        <option value="aws">AWS</option>
        <option value="azure">Azure</option>
        <option value="digital ocean">Digital Ocean</option>
        <option value="other...">Other...</option>
      </select>
      <button v-on:click="yell">Yell</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "Yell",
  data: function() {
    return {
      rHeader: '',
      pCloud: '',
      msg: '',
      cloud: '',
      response: '',
    };
  },
  watch: {
    rHeader: function(val) {
      this.rHeader = val
    },
    response: function(val) {
      this.response = val
    }
  },
  methods: {
    again: function(e) {
      e.preventDefault()

      this.rHeader = ''
      this.pCloud = ''
      this.response = ''
      this.msg = ''
      this.cloud = ''
    },
    yell: async function(e) {
      e.preventDefault()

      let response, data = null
      try {
        let encoded = {
          msg: encodeURIComponent(this.msg),
          cloud: encodeURIComponent(this.cloud),
        }
        response = await fetch(`/api/yell?msg=${encoded.msg}&cloud=${encoded.cloud}`)
        data = await response.json()

        if (response.status >= 400) {
          throw new Error()
        }
      } catch (err) {
        this.rHeader = 'Dumbass'
        this.response = data.msg
        return
      }

      // Set the response and clear the msg & cloud
      this.rHeader = 'You yelled'
      this.pCloud = this.cloud
      this.response = data.msg
      this.msg = ''
      this.cloud = ''
    }
  }
}
</script>

<style>
input[type="submit"], select {
  display: block;
  margin: 0 auto;
}

.yell {
  color: #FF0000;
}
</style>
