<template>
  <div class="main-theme">
    <div class="body-container-cover" v-if="first_enter" @click="enterPage()" align="center"><img class="mainimg" src="~/static/imgs/fengchanghome.jpg" width="1000px" alt="Click to enter!" /></div>
    <div class="body-container" v-if="!first_enter" >
      <app-header></app-header>
      <audio id="main-audio" class="sound" autoplay  preload="auto" src="http://www.???.cc/fcnotes/%E9%A3%9E-%E8%87%B4%E6%88%91%E4%BB%AC%E7%9A%84%E6%98%9F%E8%BE%B0%E5%A4%A7%E6%B5%B7.mp3" controls loop>
      </audio>
      <div style="width:100%;min-width:950px;margin-top:50px;display:flex;flex-direction:row;overflow-x:scroll;margin-left:auto;">
        <div class="duilian"></div>
        <div class="bodyimageholder">
          <a @click.prevent="goPage('/think')" href="#" class="think">思想</a>
          <a @click.prevent="goPage('/technology')" href="#" class="technology">技术</a>
          <a @click.prevent="goPage('/life')" href="#" class="life">生活</a>

        </div>

      </div>
    </div>
  </div>
</template>

<script>
import AppLogo from '~/components/AppLogo.vue'
import state from '~/common/state'
import Header from '~/components/Header.vue'
import base from '~/mixins/base'

export default {
  components: {
    AppLogo,
    'app-header': Header
  },
  watch: {
        first_enter(newValue){
            //do something if toggleElement changes
            console.log('reached');
            console.log(newValue)
            if(newValue==false){
               this.$nextTick(function () {
                   document.getElementById('main-audio').play();
              })
            }
        }
  },
  data(){
    return{
      state,
      first_enter: true,
    }
  },
  mixins:[base],
  methods:{
    goPage(val){
      this.$router.push(val)
    },
    enterPage(){
      this.first_enter = false;
      this.state.first_enter = false;
    }
  },


  mounted(){
    this.whoamifoo();
    if(this.state.first_enter==null){
       this.first_enter = true;
    }else{
      this.first_enter = this.state.first_enter;
    }
  },
  updated(){
    console.log('jump')
  }
}
</script>

<style scoped>
.main-theme {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.body-container-cover{
  width: 100%;
  height: 100%;
  background-color: black;
}
.body-container{
  width: 100%;
  height: 100%;
  background-color: white;
  position: relative;
}

.mainimg:hover{
  cursor: pointer;
  opacity: 0.8;
}






.bodyimageholder{
  position: relative;
  width: 930px;
  height: 600px;
  min-width: 930px;
  background-image: url("~/static/imgs/hand.png");
  background-size: 100% 100%;
}

.technology{
  /* background styles */
	position: absolute;
  left: 360px;
  top: 40px;
	padding: 10px 20px;
	background-color: #eaf; /*for compatibility with older browsers*/
	/* background-image: linear-gradient(green,lightgreen); */

  background: linear-gradient(#507,#eaf);
	/* text styles */
	text-decoration: none;
	color: #fff;
	font-size: 25px;
	font-family: sans-serif;
	font-weight: 100;

  border-radius: 3px;
  box-shadow: 0px 1px 4px -2px #333;
  text-shadow: 0px -1px #333;
}
.think{
  /* background styles */
	position: absolute;
  left: 100px;
  top: 95px;
	padding: 10px 20px;
	background-color: red; /*for compatibility with older browsers*/
	/* background-image: linear-gradient(green,lightgreen); */

  background: linear-gradient(#700,#f00);
	/* text styles */
	text-decoration: none;
	color: #fff;
	font-size: 25px;
	font-family: sans-serif;
	font-weight: 100;

  border-radius: 3px;
  box-shadow: 0px 1px 4px -2px #333;
  text-shadow: 0px -1px #333;
}
.life{
  /* background styles */
	position: absolute;
  left: 370px;
  top: 330px;
	padding: 10px 20px;
	background-color: orange; /*for compatibility with older browsers*/
	/* background-image: linear-gradient(green,lightgreen); */

      background: linear-gradient(#b40,orange);
	/* text styles */
	text-decoration: none;
	color: #fff;
	font-size: 25px;
	font-family: sans-serif;
	font-weight: 100;

  border-radius: 3px;
  box-shadow: 0px 1px 4px -2px #333;
  text-shadow: 0px -1px #333;
}
.think:after{
  content: '';
	position: absolute;
	top: 2px;
	left: 2px;
	width: calc(100% - 4px);
	height: 50%;
	background: linear-gradient(rgba(255,255,255,0.8), rgba(255,255,255,0.2));
}
.life:after{
  content: '';
	position: absolute;
	top: 2px;
	left: 2px;
	width: calc(100% - 4px);
	height: 50%;
	background: linear-gradient(rgba(255,255,255,0.8), rgba(255,255,255,0.2));
}
.technology:after{
  content: '';
	position: absolute;
	top: 2px;
	left: 2px;
	width: calc(100% - 4px);
	height: 50%;
	background: linear-gradient(rgba(255,255,255,0.8), rgba(255,255,255,0.2));
}
.think:hover{
	background: linear-gradient(rgb(134, 18, 43),rgb(134, 108, 111));
}
.technology:hover{
	background: linear-gradient(rgb(10, 131, 168),rgb(129, 135, 136));
}
.life:hover{
	background: linear-gradient(rgb(165, 168, 10),rgb(159, 165, 144));
}
.duilian{
  display: block;
  height: 600px;
  width: 200px;
  min-width: 200px;
  margin-left: auto;
  background-image: url("~/static/imgs/pcwn.png");
  background-size: 100% 100%;
}
.sound{
  width: 280px;
  height:40px;
  top:85px;
  left:700px;
  position:absolute;
}
</style>
