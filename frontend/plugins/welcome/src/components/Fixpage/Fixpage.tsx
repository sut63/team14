//style
import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField,Button } from '@material-ui/core';
import InputAdornment from '@material-ui/core/InputAdornment';
import { Cookies } from '../WelcomePage/Cookie'
import { DefaultApi } from '../../api/apis';
//icon
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import PersonIcon from '@material-ui/icons/Person';
import ComputerTwoToneIcon from '@material-ui/icons/ComputerTwoTone';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
//alert
import Swal from 'sweetalert2';
//entity
import { EntPersonal } from '../../api/models/EntPersonal';
import { EntFixcomtype } from '../../api/models/EntFixcomtype';
import { EntCustomer } from '../../api/models/EntCustomer';
import { EntFixbrand } from '../../api/models/EntFixbrand';
import { EntFix } from '../../api/models/EntFix';

// css style
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

//alert setting
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

//Cookies
var ck = new Cookies()
  var cookieName = ck.GetCookie()
  var cookieID = ck.GetID()

export default function Fixpage() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [fixs, setFixs] = useState<EntFix[]>([]);
  const [productnumber, setProductnumber] = useState(String);
  const [date, setDate] = useState(String);
  const [problemtype, setProblemtype] = useState(String);
  const [queue, setQueue] = useState(String);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  const [personal, setPersonal] = useState(Number);

  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);
  const [customer, setCustomer] = useState(Number);

  const [fixcomtypes, setFixcomtypes] = React.useState<EntFixcomtype[]>([]);
  const [fixcomtype, setFixcomtype] = useState(Number);

  const [fixbrands, setFixbrands] = React.useState<EntFixbrand[]>([]);
  const [fixbrand, setFixbrand] = useState(Number);

//validate
  const [productnumberError, setProductnumberError] = React.useState('');
  const [problemtypeError, setProblemtypeError] = React.useState('');
  const [queueError, setQueueError] = React.useState('');
  
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getPersonals = async () => {
      const res = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();

    const getCustomers = async () => {
      const res = await http.listCustomer({ limit: 10, offset: 0 });
      setLoading(false);
      setCustomers(res);
      console.log(res);
    };
    getCustomers();

    const getFixcomtypes = async () => {
      const res = await http.listFixcomtype({ limit: 10, offset: 0 });
      setLoading(false);
      setFixcomtypes(res);
      console.log(res);
    };
    getFixcomtypes();

    const getFixbrands = async () => {
        const res = await http.listFixbrand({ limit: 10, offset: 0 });
        setLoading(false);
        setFixbrands(res);
        console.log(res);
      };
      getFixbrands();

  }, [loading]);

  const getFix = async () => {
    const res = await http.listFix({ limit: 10, offset: 0 });
    setFixs(res);
  };

  const handleProductnumberChange = (event: React.ChangeEvent<{value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('Productnumber', validateValue)
    setProductnumber(event.target.value as string);
  };

  const handleDateChange = (event: any) => {
    setDate(event.target.value as string);
  };

  const handleProblemtypeChange = (event: React.ChangeEvent<{value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('Problemtype', validateValue)
    setProblemtype(event.target.value as string);
  };

  const handleQueueChange = (event: React.ChangeEvent<{value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('Queue', validateValue)
    setQueue(event.target.value as string);
  };
  
  const CustomerhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCustomer(event.target.value as number);
  };

  const PersonalhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonal(event.target.value as number);
  };

  const FixcomtypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFixcomtype(event.target.value as number);
  };

  const FixbrandhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFixbrand(event.target.value as number);
  };

  //let p = Number(personal)

  const fix = {
    productnumber : productnumber,
    date : date + ":00+07:00",
    problemtype : problemtype,
    queue : queue,
    personal : Number(cookieID),
    customer : customer,
    fixcomtype : fixcomtype,
    fixbrand : fixbrand,
  };

// ฟังก์ชั่นสำหรับ validate หมายเลขผลิตภัณฑ์
  const validareProductnumber = (val: string)=>{
    return val;
  }
// ฟังก์ชั่นสำหรับ validate รายละเอียดการแจ้งซ่อม
  const validareProblemtype = (val: string)=>{
    return val.length >=5 && val.length <=100 ? true : false;
  }
// ฟังก์ชั่นสำหรับ validate ลำดับคิว
  const validareQueue = (val: string)=>{
    return val.match("[F]+[I]+[X]+[-]") && val.length == 7;
  }

 // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
 const checkPattern  = (id: string, value: string) => {
  switch(id) {
    case 'Productnumber':
      validareProductnumber(value) ? setProductnumberError('') : setProductnumberError('กรอกหมายเลขผลิตภัณฑ์ a-z,A-Z,0-9');
      return;
    case 'Problemtype':
      validareProblemtype(value) ? setProblemtypeError('') : setProblemtypeError('กรอกรายละเอียดการแจ้งซ่อม 5 ตัวอักษรขึ้นไป');
      return;
    case 'Queue':
      validareQueue(value) ? setQueueError('') : setQueueError('รูปแบบหมายเลข FIX-XXX')
      return;
    default:
      return;
  }
}
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
  
  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'Productnumber':
        alertMessage("error","กรุณากรอกหมายเลขผลิตภัณฑ์");
        return;
      case 'Problemtype':
        alertMessage("error","กรุณากรอกรายละเอียดการแจ้งซ่อม");
        return;
        case 'Queue':
          alertMessage("error","กรุณากรอกลำดับคิว");
          return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  const save = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/fixs';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(fix),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  };



  /*// create fix
const CreateFix = async () => {
  if ((personal != null) && (customer != null) && (fixcomtype != null) && (fixbrand != null) && (productnumber != "")
      && (date != "") && (problemtype != "") && (queue != "")){
  const fix = {
    productnumber : productnumber,
    date : date + ":00+07:00",
    problemtype : problemtype,
    queue : queue,
    personal : Number(cookieID),
    customer : customer,
    fixcomtype : fixcomtype,
    fixbrand : fixbrand,
  };
  console.log(fix);
  const res: any = await http.createFix({ fix: fix });
 
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
       title="ระบบบันทึกการแจ้งซ่อมคอมพิวเตอร์" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button onClick={() => {save();}} variant="contained"  color="primary" startIcon={<AddCircleOutlinedIcon/>}> เพิ่มข้อมูลการแจ้งซ่อมสินค้า </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button style={{ marginLeft: 20 }} component={RouterLink} to="/Tablefix" variant="contained" startIcon={<CancelRoundedIcon/>}> ดูข้อมูล </Button>
     </Header>

     <Content>
     <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
          >
           
            <div className={classes.paper}><strong>ชื่อลูกค้า</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              labelId="customername-label"
              id="customername"
              value={customer}
              onChange={CustomerhandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">ชื่อลูกค้า</InputLabel>

              {customers.map((item: EntCustomer) => (
                <MenuItem value={item.id}>{item.customername}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>ประเภทคอมพิวเตอร์</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              id="fixcomtype"
              value={fixcomtype}
              onChange={FixcomtypehandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกประเภทคอมพิวเตอร์</InputLabel>

              {fixcomtypes.map((item: EntFixcomtype) => (
                <MenuItem value={item.id}>{item.fixcomtypename}</MenuItem>
              ))}
            </Select>

            
            <div className={classes.paper}><strong>แบรนด์</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              labelId="fixbrandname-label"
              id="fixbrandname"
              value={fixbrand}
              onChange={FixbrandhandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">เลือกแบรนด์คอมพิวเตอร์</InputLabel>

              {fixbrands.map((item: EntFixbrand) => (
                <MenuItem value={item.id}>{item.fixbrandname}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>หมายเลขผลิตภัณฑ์</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <ComputerTwoToneIcon />
                </InputAdornment>
              ),
            }}
            error = {productnumberError ? true : false}
              id="productnumber"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              helperText= {productnumberError}
              value={productnumber}
              onChange={handleProductnumberChange}
            />

            <div className={classes.paper}><strong>รายละเอียดการแจ้งซ่อม/ปัญหา</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <ComputerTwoToneIcon />
                </InputAdornment>
              ),
            }}
            error = {problemtypeError ? true : false}
              id="problemtype"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              helperText= {problemtypeError}
              value={problemtype}
              onChange={handleProblemtypeChange}
            />

            <div className={classes.paper}><strong>ลำดับคิว</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <PersonIcon />
                </InputAdornment>
              ),
            }}
            error = {queueError ? true : false}
              id="queue"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              helperText= {queueError}
              value={queue}
              onChange={handleQueueChange}
            />
          
          </FormControl>

        </form>
      </div>
      </Content>
   </Page>
 );
}
