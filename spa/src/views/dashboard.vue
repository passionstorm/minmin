<template>
  <div>
    <div class="cards">
      <card>saccasjj</card>
    </div>
    <button type="button">Click here</button>
  </div>
</template>
<script>
  export default {
    components: {
      card: () => import('_w/Card'),
    },
    methods: {},
    data() {
      return {
        abc: 'ngu',
      };
    },
    computed: {},
    created() {
      // this.$emit("slot-footer", () => import("_c/SlotContent"));
      // bus.$on("dashboard.click_save", data => {
      // });
    },
    mounted: function() {
    },
    destroyed: function() {
      // bus.$off("dashboard.click_save");
      // this.$emit("slot-footer", () => import("_c/SlotContent"));
    },
  };
</script>

<style scoped>
</style>