<template>
  <div>
    <button
      class="btn"
      @click="test"
    >click</button>
    <button>back</button>
  </div>
</template>

<script>
export default {
  methods: {
    test() {
        console.log(this.$bus);
        
      this.$bus.$emit("dashboard.click_save", "me");
    }
  }
};
</script>

<style>
</style>