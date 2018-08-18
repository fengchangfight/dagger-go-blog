<template>
  <div class="about-page">
      <div style="width:100%;height:100%;">
          <app-header></app-header>
          <div class="about-me-container">
            <div class="back-music-fcblog">
                   <audio style="width: 320px;height:40px;" src="http://www.???.cc/fcnotes/%E9%AA%84%E5%82%B2%E7%9A%84%E5%B0%91%E5%B9%B4.mp3" controls autoplay loop>
                   </audio>
                </div>
            <div class="big-profile"></div>
            <div class="about-me-content">
              <div class="contact-group">
                <a class="contact" style="text-decoration:none;" href="mailto:???@gmail.com"><i class="fa fa-envelope" aria-hidden="true"></i>&nbsp;Gmail</a>
                <a class="contact" href="https://github.com/???"><i class="fa fa-github" aria-hidden="true"></i>&nbsp;Github</a>
                <a class="contact" href="https://www.linkedin.com/in/????/"><i class="fa fa-linkedin" aria-hidden="true"></i>&nbsp;LinkedIn</a>
                <a class="contact" href="https://twitter.com/????"><i class="fa fa-twitter" aria-hidden="true"></i>&nbsp;Twitter</a>
              </div>
            </div>
          </div>
      </div>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
export default {
  components: {
      'app-header': Header
  },
  data(){
    return{

    }
  }
}
</script>

<style scoped>
.about-page {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.about-me-container{
  width: 900px;
  min-width: 850px;
  min-height: 600px;
  height: 600px;
  margin:0 auto;
  margin-top:10px;
  display: flex;
  flex-direction: column;
}
.big-profile{
  height: 400px;
  width: 100%;
  background-color: gray;
  background-image: url('~/static/imgs/programmer.jpeg');
  background-size: cover;
}
.about-me-content{
  flex: 1;
  background-color: white;
  height: 100%;
}
.contact{
   text-decoration:none;
   margin-left: 20px;
}
.contact-group{
  margin-top: 10px;
}
.back-music-fcblog{
  margin-top: 20px;
}
</style>
