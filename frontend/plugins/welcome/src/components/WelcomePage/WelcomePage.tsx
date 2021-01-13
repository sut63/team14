import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { TextField } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi, EntPersonal } from '../../api';
import { Alert } from '@material-ui/lab';

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

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);
  const [loading, setLoading] = useState(false);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");

  useEffect(() => {
    const getPersonals = async () => {
      const res:any = await http.listPersonal({});
      setLoading(false);
      setPersonals(res);
    }
    getPersonals();

    const resetPersonalData = async () => {
      setLoading(false);
      localStorage.setItem("userdata", JSON.stringify(null));
    }
    resetPersonalData();

  }, [loading]);

  const HandleEmailChange = (event:any) => {
    setEmail(event.target.value as string);
  };

  const HandlePasswordChange = (event:any) => {
    setPassword(event.target.value as string);
  };

  const Checklogin = async () => {
    personals.map((item: any) => {
      console.log(item.email);
      if ((item.email == email) && (item.password == password)) {
        setAlert(true);
        localStorage.setItem("personaldata", JSON.stringify(item.id));
        history.pushState("", "", "/Group14");
        window.location.reload(false);
      }
    })
    setStatus(true);
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  }

  return (
    <Page theme={pageTheme.tool}>
    <Header
      title="ระบบแจ้งซ่อมคอมพิวเตอร์" type="กลุ่มที่14">
    </Header>
    <Content>
      <ContentHeader title="เข้าสู่ระบบการแจ้งซ่อม">
      {status ? (
            <div>
              {alert ? (
                <Alert severity="success" onClose={() => { }}>
                  เข้าสู่ระบบสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="error" onClose={() => { setStatus(false) }}>
                    เข้าสู่ระบบไม่สำเร็จ กรุณาตรวจสอบ Email หรือ Password
                  </Alert>
                )}
            </div>
          ) : null}
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
            value={email}
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
            value={password}
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
           onClick={() => {
              Checklogin();
            }}
          >
            เข้าสู่ระบบ
          </Button>    
      </div>
      </Content>
   </Page>
  );
}