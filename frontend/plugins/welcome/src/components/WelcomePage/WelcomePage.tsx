import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { TextField } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi, EntPersonal } from '../../api';
import { Alert } from '@material-ui/lab';
import { Cookies } from './Cookie';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  margin: {
    width: 500 ,
     marginLeft:7,
     marginRight:-7,
  },
  submit: {
    background: 'linear-gradient(45deg, #20002c 30%, #cbb4d4 90%)',
    color: 'white',
  },
  }),
);

export default function SignInSide() {
  const classes = useStyles();
  const http = new DefaultApi();
  var ck = new Cookies();
  var check : boolean
  const [path, setPath] = React.useState("");

  const [status, setStatus] = useState(false);
  //const [alert, setAlert] = useState(Boolean);
  const [loading, setLoading] = useState(false);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");

  const getPersonals = async () => {
    const res = await http.listPersonal({});
    setPersonals(res);
  }
  getPersonals();

  const HandleEmailChange = (event:any) => {
    setEmail(event.target.value as string);
  };
  const HandlePasswordChange = (event:any) => {
    setPassword(event.target.value as string);
  };

  useEffect(() => {
    getPersonals();
  }, []);

  const Checklogin = async () => {
    check = ck.CheckLogin(personals,email,password)
    console.log("check => "+check)
    if(check === true){
      history.pushState('', '', '/Group14');
      ck.SetCookie("email",email,30)
      ck.SetCookie("id",ck.SetID(personals,email,password),30)
      ck.SetCookie("name",ck.SetName(personals,email,password),30)
      window.location.reload(false)
      alert("เข้าสู่ระบบสำเร็จ")
    }
    else if(check === false){
      alert("เข้าสู่ระบบไม่สำเร็จ กรุณาตรวจสอบ Email หรือ Password")
      //setPath("/")
    }
  }

  return (
    <Page theme={pageTheme.tool}>
    <Header
      title="ระบบแจ้งซ่อมคอมพิวเตอร์" type="กลุ่มที่14">
    </Header>
    <Content>
    <ContentHeader title="เข้าสู่ระบบการแจ้งซ่อม">
    </ContentHeader>
      <div className={classes.root}>
      <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            onChange={HandleEmailChange}
            autoComplete="email"
            autoFocus
            style={{ width: 400 , marginRight: 400, marginLeft: 400}}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            onChange={HandlePasswordChange}
            autoComplete="current-password"
            style={{ width: 400 , marginRight: 400, marginLeft: 400}}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            size="large"
            className={classes.submit}
            style={{ width: 400 , marginRight: 400, marginLeft: 400, marginTop: 2}}
            onClick={() => {Checklogin();}}
          >
            เข้าสู่ระบบ
          </Button>    
      </div>
      </Content>
   </Page>
  );
}