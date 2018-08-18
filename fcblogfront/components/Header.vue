<template>
  <div class="nav-container">
         <nav style="width:100%;position:relative;z-index:1;">
           <ul style="">
             <li><a href="" @click.prevent="goPage('/')" >首页</a></li>
             <li><a href="" @click.prevent="goPage('/allposts')">所有文章</a></li>
             <li><a href="" @click.prevent="goPage('/pictures')">图库</a></li>
             <li><a href="" @click.prevent="goPage('/about')">关于我</a></li>
           </ul>
            <div v-if="whoami!=null&&whoami.length>0" class="dropdown" style="position:absolute;left:800px;top:14px;z-index:100;">
                 <button class="dropbtn">{{whoami}}<label style="color:red;" v-if="Number(unread_comment_cnt)>0">({{unread_comment_cnt}})</label></button>
                 <div class="dropdown-content" style="z-index:100;" >
                   <a href="#" @click.prevent="goPage('/comment-admin')" style="z-index:100;">评论管理<label style="color:red;" v-if="Number(unread_comment_cnt)>0">({{unread_comment_cnt}})</label></a>
                   <a href="#" @click.prevent="logout" style="z-index:100;">登出</a>
                 </div>
            </div> 
         </nav>
        
   </div>
</template>

<script>
import {AXIOS} from '~/common/http-commons'
import state from '~/common/state'
import base from '~/mixins/base'

export default {
  name: 'Header',
  data() {
      return {
          state,
      };
  },
  mixins:[base],
  components: {
  },
  methods: {
    logout () {
        AXIOS.post('/logout', {}, {headers:{'Content-Type':'application/x-www-form-urlencoded'}})
          .then(response => {
            window.location.href = '/';
          })
          .catch(e => {
            window.location.href = '/';
          })
      },
    goPage(val){
        this.state.first_enter=false;
        this.$router.push(val)
    },
  },
  mounted () {
    this.whoamifoo();
    this.countUreadComment();
  },
  created(){

  },
  updated() {

  }
}
</script>

<style scoped>
nav{
  background: rgba(70, 66, 66, 0.5);
  box-shadow: 0px 7px 7px rgb(61, 58, 58);
  overflow: visible;
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
li.right{
  float: right;
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
  position: relative;
}
nav a:hover{
  background: rgba(58, 42, 42, 0.6);
  border-radius: 15px 50px 30px;
}
.nav-container{
  max-width: 1000px;
  min-width: 900px;
  margin: 0 auto;
  margin-top: 10px;
  display: flex;
  flex-direction: row;
}

.dropbtn {
    background-color: #a2a0a0;
    color: white;
    padding: 3px;
    margin-right: 100px;
    margin-top: 2px;
    font-size: 16px;
    border: none;
    cursor: pointer;
}
.dropbtn:hover{
  opacity: 0.5;
}

/* The container <div> - needed to position the dropdown content */
.dropdown {
    position: relative;

}

/* Dropdown Content (Hidden by Default) */
.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
}

/* Links inside the dropdown */
.dropdown-content a {
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    
}

/* Change color of dropdown links on hover */
.dropdown-content a:hover {background-color: #f1f1f1}

/* Show the dropdown menu on hover */
.dropdown:hover .dropdown-content {
    display: block;
}

/* Change the background color of the dropdown button when the dropdown content is shown */
.dropdown:hover .dropbtn {
    background-color: #a2a0a0;
}
</style>
