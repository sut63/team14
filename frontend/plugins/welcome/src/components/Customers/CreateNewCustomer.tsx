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
//import { EntPersonal, EntGender, EntTitle, EntCustomer} from '../../api';
import { Cookies } from '../WelcomePage/Cookie';
import { EntPersonal } from '../../api/models/EntPersonal';
import { EntGender } from '../../api/models/EntGender';
import { EntTitle } from '../../api/models/EntTitle';
import { EntCustomer } from '../../api/models/EntCustomer';
import Swal from 'sweetalert2';
import Customers from '.';

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

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

var ck = new Cookies()
var cookieName = ck.GetCookie()
var cookieID = ck.GetID()

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
  const [identificationnumber, setIdentificationnumber] = useState(String);


  useEffect(() => {
    const getPersonals = async () => {
      const p = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(p);
      
    };
    getPersonals();

    const getTiles = async () => {
      const t = await http.listTitle({ limit: 10, offset: 0 });
      setLoading(false);
      setTitles(t);
      
    };
    getTiles();

    const getGenders = async () => {
      const g = await http.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(g);
      
    };
    getGenders();

  }, [loading]);

  /*const getCustomer = async () => {
    const res = await http.listCustomer({ limit: 10, offset: 0 });
    setCustomer(res);
  };*/

  const handleCustomernameChange = (event: any) => {
    setCustomername(event.target.value as string);
  };

  const handleAddressChange = (event: any) => {
    setAddress(event.target.value as string);
  };

  const handlePhonenumberChange = (event: any) => {
    setPhonenumber(event.target.value as string);
  };

  const handleIdentificationnumberChange = (event: any) => {
    setIdentificationnumber(event.target.value as string);
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

 let p = Number(personal)

 const customer = {
  customername : customername,
  address : address,
  phonenumber : phonenumber,
  identificationnumber : identificationnumber,
  title : title,
  personal : Number(cookieID),
  gender : gender,
  };

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'Customername':
        alertMessage("error","กรุณากรอกชื่อ");
        return;
      case 'Address':
        alertMessage("error","กรุณากรอกที่อยู่");
        return;
      case 'Phonenumber':
        alertMessage("error","กรุณากรอกข้อมูลหมายเลขโทรศัพท์ 10 หลัก");
        return;
      case 'Identificationnumber':
        alertMessage("error","กรุณากรอกข้อมูลบัตรประชาชน 13 หลัก");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  console.log(customer)
function save() {
  const apiUrl = 'http://localhost:8080/api/v1/customers';
  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(customer),
  };

  console.log(customer); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

  fetch(apiUrl, requestOptions)
    .then(response => response.json())
    .then(data => {
      console.log(data);
      if (data.status == true) {
        //clear();
        Toast.fire({
          icon: 'success',
          title: 'บันทึกข้อมูลสำเร็จ',

        });//window.setTimeout(function(){location.reload()},8000);
      } else {
        checkCaseSaveError(data.error.Name)
      }
    });

    
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };
  

  // create personal
/*const CreateCustomer = async () => {
  if ((customername != null) && (customername != "") && (address != null) && (address != "") && (phonenumber != null) && (phonenumber != "") && (title != null) && (personal != null) && (gender != null) ) {
  
    const customer = {
    customername : customername,
    address : address,
    phonenumber : phonenumber,
    title : title,
    personal : Number(cookieID),
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
};*/

 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="Customer System" type="Computer Repair System">
     </Header>

     <Content>
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

<           div className={classes.paper}><strong>หมายเลขบัตรประชาชน</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <PersonIcon />
                </InputAdornment>
              ),
            }}
              id="identificationnumber"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={identificationnumber}
              onChange={handleIdentificationnumberChange}
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

          
          </FormControl>
        
       <div>&nbsp;&nbsp;&nbsp;</div>
      <Button onClick={() => {save();}} variant="contained" color="primary" startIcon={<AddCircleOutlinedIcon/>}> Create new customer </Button>
      <Button style={{ marginLeft: 20 }} component={RouterLink} to="/customertable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Table </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
     
     
    

        </form>
      </div>
      </Content>
   </Page>
 );
}