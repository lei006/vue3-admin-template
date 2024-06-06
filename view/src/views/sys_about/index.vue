<template>
  <div class="layout-container">

    <div class="container">
      <div class="title">
        <h1 class="font-normal text-center py-8">About Northwoods</h1>
      </div>

      <div class="cont-row">
          <div class="text">
            <p>
              Northwoods Software Corporation is a privately-held company, founded in 1995 by a group of engineers with a shared vision for providing superior graphical user interfaces.
              We have been creating diagram components for over 20 years, over many platforms, and reap the benefits of redesigning a library from scratch for each new platform while taking into account our accumulated experience.
            </p>
            <p>
              We have grown to become a world-class supplier of interactive diagram components and class libraries across a variety of platforms.
            </p>
            <p>
              We are a company created by engineers for engineers. We pride ourselves on quality software design and excellent developer-to-developer technical support.
            </p>
          </div>
          <div class="image">
            <img class="rounded shadow-md" src="https://nwoods.com/assets/images/waterStreet.jpeg">
            <p class="text-sm text-gray-500 text-center">4 Water Street, Nashua, New Hampshire</p>
          </div>
      </div>
    </div>

  </div>
</template>

<script setup>

</script>

<style lang="scss" scoped>
  .container {
    margin: auto auto;
    width: 1000px;
    display: flex;
    flex-direction: column;
  }

  .title {
    text-align: center;
  }


  .cont-row {
    display: flex;
    flex-direction: row;
  }

  .cont-row .text {
    flex:1;
    padding: 10px;
    font-size: 16px;
  }

  .cont-row .image {
    flex:1;
    padding: 10px;    

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .shadow-md {
    width: 350px;
    height: 350px;
  }

</style>
