<template>
  <div class="card"><slot></slot></div>
</template>

<script>
export default {};
</script>

<style scoped>
.card {
  position: relative;
  display: -ms-flexbox;
  display: flex;
  -ms-flex-direction: column;
  flex-direction: column;
  min-width: 0;
  word-wrap: break-word;
  background-color: #fff;
  background-clip: border-box;
  box-shadow: 0 0 1px rgba(0, 0, 0, 0.125), 0 1px 3px rgba(0, 0, 0, 0.2);
  border: 0 solid rgba(0, 0, 0, 0.125);
  border-radius: 0.25rem;
  padding: 1.25rem;
}
</style>>
