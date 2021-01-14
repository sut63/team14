import React, {  FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField,Button } from '@material-ui/core';
import { Alert, AlertTitle } from '@material-ui/lab';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntPersonal, EntGender, EntTitle, EntCustomer} from '../../api';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  margin: {
     margin: theme.spacing(2),
  },
  insideLabel: {
   margin: theme.spacing(1),
 },
  button: {
   marginLeft: '40px',
 },
  textField: {
   width: 500 ,
   marginLeft:7,
   marginRight:-7,
  },
   select: {
     width: 500 ,
     marginLeft:7,
     marginRight:-7,
     //marginTop:10,
   },
   paper: {
     marginTop: theme.spacing(1),
     marginBottom: theme.spacing(1),
     marginLeft: theme.spacing(1),
   },
   pa: {
     marginTop: theme.spacing(2),
   },
  }),
);

export default function CreateNewCustomer() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [customers, setCustomer] = React.useState<EntCustomer[]>([]);
  const [titles, setTitles] = React.useState<EntTitle[]>([]);
  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);
  const [loading, setLoading] = useState(true);

  const [personal, setPersonal] = useState(Number);
  const [title, setTitle] = useState(Number);
  const [gender, setGender] = useState(Number);

  const [customername, setCustomername] = useState(String);
  const [address, setAddress] = useState(String);
  const [phonenumber, setPhonenumber] = useState(String);

  useEffect(() => {
    const getPersonals = async () => {
      const res = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();

    const getTiles = async () => {
      const res = await http.listTitle({ limit: 10, offset: 0 });
      setLoading(false);
      setTitles(res);
      console.log(res);
    };
    getTiles();

    const getGenders = async () => {
      const res = await http.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(res);
      console.log(res);
    };
    getGenders();

  }, [loading]);

  const getCustomer = async () => {
    const res = await http.listCustomer({ limit: 10, offset: 0 });
    setCustomer(res);
  };

  const handleCustomernameChange = (event: any) => {
    setCustomername(event.target.value as string);
  };

  const handleAddressChange = (event: any) => {
    setAddress(event.target.value as string);
  };

  const handlePhonenumberChange = (event: any) => {
    setPhonenumber(event.target.value as string);
  };
  
  const TitlehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTitle(event.target.value as number);
  };

  const PersonalhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonal(event.target.value as number);
  };

  const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGender(event.target.value as number);
  };

 
  // create personal
const CreateCustomer = async () => {
  if ((customername != null) && (customername != "") && (address != null) && (address != "") && (phonenumber != null) && (phonenumber != "") && (title != null) && (personal != null)  && (gender != null) ) {
  
    const customer = {
    customername : customername,
    address : address,
    phonenumber : phonenumber,
    title : title,
    personal : personal,
    gender : gender,
    };
  console.log(customers);
  const res: any = await http.createCustomer({ customer: customer });

    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } 
  }
    else {
      setStatus(true);
      setAlert(false);
    }
  const timer = setTimeout(() => {
    setStatus(false);
  }, 3000);
};

 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="Customer System" type="Computer Repair System">
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button onClick={() => {CreateCustomer();}} variant="contained" href="/customertable" color="primary" startIcon={<AddCircleOutlinedIcon/>}> Create new customer </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
     
      <Button style={{ marginLeft: 20 }} component={RouterLink} to="/customertable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Table </Button>
     </Header>

     <Content>
       {status ? ( 
        <div>
        {alert ? ( 
            <Alert severity="success"  onClose={() => { }}> 
              <AlertTitle> บันทึกข้อมูลสำเร็จ </AlertTitle></Alert>) 
      : (     
        <Alert severity="error" onClose={() => { setStatus(false) }}> 
          <AlertTitle> ไม่สามารถบันทึกข้อมูลได้ กรุณาลองใหม่อีกครั้ง </AlertTitle></Alert>)}
      </div>
          ) : null}
     <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
          >
            <div className={classes.paper}><strong>คำนำหน้าชื่อ</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              labelId="nametitle-label"
              id="nametitle"
              value={title}
              onChange={TitlehandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">เลือกคำนำหน้าชื่อ</InputLabel>

              {titles.map((item: EntTitle) => (
                <MenuItem value={item.id}>{item.titlename}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>ชื่อ-นามสกุล</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <PersonIcon />
                </InputAdornment>
              ),
            }}
              id="customername"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={customername}
              onChange={handleCustomernameChange}
            />

            <div className={classes.paper}><strong>ที่อยู่</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  
                </InputAdornment>
              ),
            }}
              id="address"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={address}
              onChange={handleAddressChange}
            />

            <div className={classes.paper}><strong>เบอร์โทร</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  
                </InputAdornment>
              ),
            }}
              id="phonenumber"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={phonenumber}
              onChange={handlePhonenumberChange}
            />
            
            <div className={classes.paper}><strong>เพศ</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              id="gender"
              value={gender}
              onChange={GenderhandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกเพศ</InputLabel>

              {genders.map((item: EntGender) => (
                <MenuItem value={item.id}>{item.gendername}</MenuItem>
              ))}
            </Select>
            
            <div className={classes.paper}><strong>เจ้าหน้าที่</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              id="personal"
              value={personal}
              onChange={PersonalhandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกเจ้าหน้าที่</InputLabel>

              {personals.map((item: EntPersonal) => (
                <MenuItem value={item.id}>{item.personalname}</MenuItem>
              ))}
            </Select>

          
          </FormControl>
          
          <Button style={{ marginLeft: 20 }} component={RouterLink} to="/welcome" variant="contained" >  BACK </Button>

        </form>
      </div>
      </Content>
   </Page>
 );
}