<template>
  <div class="admin-login-page">
    <div class="login-form" v-loading="wholepageloading">
          <div class="input-grouper">
              <el-input style="" v-model="username" placeholder="用户名"></el-input>
              <el-input style="margin-top: 17px;" v-model="password" type="password" placeholder="密码"></el-input>
              <a class="si-login green-button" @click="login" >登录</a>
          </div>

      </div>
  </div>
</template>

<script>
import {AXIOS} from '~/common/http-commons'
import Qs from 'qs'
import Header from '~/components/Header.vue'
import swal from 'sweetalert'
import base from '~/mixins/base'
export default {
  components: {
      'app-header': Header
  },
  mixins:[base],
  methods:{
      goPage(val){
          this.$router.push(val);
      },
      checklogin () {
        AXIOS.get('/api/v1/user/whoami').then(response => {
              // JSON responses are automatically parsed.
              this.whoami = response.data;
              // if already logged in, go to home page
              this.goPage('/')
            }).catch(e => {
              
            })
      },
      login () {
        if(this.username==null || this.username.trim().length<1 || this.password==null || this.password.trim().length<1){
              swal ( "提示" ,  "用户名或密码不能为空" ,  "info" );
              return;
        }
        this.wholepageloading = true;
        var data = {'username': this.username, 'password':this.password, 'timeout':1000};
        AXIOS.post('/login', Qs.stringify(data), {headers:{'Content-Type':'application/x-www-form-urlencoded'}}).then(response => {
            window.location.href = '/';
        }).catch(e => {
          this.wholepageloading = false;
          this.$notify.error({
            title: '错误',
            message: '用户名或密码错误'
          });
          this.$router.push('/')
          console.log(e)
        })
      }
  },
  mounted(){
    this.checklogin();
  },
  data(){
    return{
        wholepageloading: false,
        username: '',
        password: ''
    }
  }
}
</script>

<style scoped>
.admin-login-page {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;

}
nav{
  background: rgba(70, 66, 66, 0.5);
  box-shadow: 0px 7px 7px rgb(61, 58, 58);
  overflow: auto;
  width: 100;
  border-radius: 15px 50px 30px;
}
ul{
  list-style: none;
  margin: 0;
  padding: 0;
}
li{
  float: left;
}
nav a{
  width: 120px;
  display: block;
  padding: 15px 15px;
  font-family: HYDaSongJ;
  font-size: 17px;
  color: white;
  text-align: center;
  text-decoration: none;
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
}

.mainimg:hover{
  cursor: pointer;
  opacity: 0.8;
}

.nav-container{
  max-width: 1000px;
  margin: 0 auto;
  margin-top: 10px;
}

nav a:hover{
  background: rgba(58, 42, 42, 0.6);
}


.login-form{
    position: absolute;
    width: 380px;
    height: 230px;
    background-color: rgba(181, 164, 204, 0.5);
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
    -moz-box-shadow: 1px 1px 25px #202020; /* 老的 Firefox */
    box-shadow: 1px 1px 25px #202020;
    border-radius: 20px;
}
.input-grouper{
    width: 80%;
    height: 80%;
    /* background-color: aliceblue; */
    margin: 0 auto;
    margin-top: 20px;
    display: block;
}
.si-login,.si-login:visited{
    padding:5px 132px;
    border-radius: 10px;
    font-size: 19px;
    text-decoration: none;
    color: white !important;
    display: inline-block;
    margin: 0 auto;
    margin-top: 27px;
}

.green-button{
    background-color: rgb(39, 199, 33);
}
.green-button:hover{
    cursor: pointer;
    background-color: rgba(44, 33, 199,0.5);
    transform: translateY(-2px);
}


</style>

