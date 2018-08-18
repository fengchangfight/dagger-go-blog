import {AXIOS} from '~/common/http-commons'

export default{
  methods :{
    countUreadComment(){
      AXIOS.get("/api/v1/comment/unreadcnt").then(response=>{
        this.unread_comment_cnt = response.data;
        console.log(this.unread_comment_cnt);
      }).catch(e=>{

      })
    },
    whoamifoo () {
        AXIOS.get('/api/v1/user/whoami').then(response => {
              // JSON responses are automatically parsed.
              this.whoami = response.data;
            }).catch(e => {
              
            })
    },
    checkNotEmpty(st){
        if(st!=null && st.toString().length>0){
            return true;
        }else{
            return false;
        }
    },
    goPage (val) {
      this.$router.push(val)
    },
    goPost(id){
      this.$router.push('/post/'+id);
    },
    editBlog(id){
      this.$router.push('/editpost/'+id);
    },
    trashBlog(id){
      this.$confirm('确认移到垃圾箱？')
      .then(_ => {
          AXIOS.put('/api/v1/trashblog/'+id).then(response=>{
              if(response.data.ok==true){
                  this.$notify({
                     title: '成功',
                     type: 'success',
                     message: response.data.message
                   });
                   this.loadPosts();
              }else{
                  this.$notify.error({
                     title: '错误',
                     message: response.data.message
                   });
              }
          }).catch(e=>{
              this.$notify.error({
                     title: '错误',
                     message: '未知错误'
                   });
          })
          done();
      }).catch(_ => {});
 },
  },
  data () {
    return {
      unread_comment_cnt:'',
      whoami:'',
      head_img:'head.png'
    }
  }
}
