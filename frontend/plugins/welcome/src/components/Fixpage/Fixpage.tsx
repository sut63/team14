import React, {  FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField,Button } from '@material-ui/core';
import { Alert, AlertTitle } from '@material-ui/lab';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import ComputerTwoToneIcon from '@material-ui/icons/ComputerTwoTone';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntPersonal, EntFixcomtype, EntCustomer, EntBrand, EntFix} from '../../api';

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

  const [brands, setBrands] = React.useState<EntBrand[]>([]);
  const [brand, setBrand] = useState(Number);
  
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

    const getBrands = async () => {
        const res = await http.listBrand({ limit: 10, offset: 0 });
        setLoading(false);
        setBrands(res);
        console.log(res);
      };
      getBrands();

  }, [loading]);

  const getFix = async () => {
    const res = await http.listFix({ limit: 10, offset: 0 });
    setFixs(res);
  };

  const handleProductnumberChange = (event: any) => {
    setProductnumber(event.target.value as string);
  };

  const handleDateChange = (event: any) => {
    setDate(event.target.value as string);
  };

  const handleProblemtypeChange = (event: any) => {
    setProblemtype(event.target.value as string);
  };

  const handleQueueChange = (event: any) => {
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

  const BrandhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBrand(event.target.value as number);
  };

  // create fix
const CreateFix = async () => {
  const fix = {
    productnumber : productnumber,
    date : date + ":00+07:00",
    problemtype : problemtype,
    queue : queue,
    personal : personal,
    customer : customer,
    fixcomtype : fixcomtype,
    brand : brand,
  };
  console.log(fixs);
  const res: any = await http.createFix({ fix: fix });
  console.log("hi");
  setStatus(true);
  
  if (res.id != '') {
    setAlert(true);
  } else {
    setAlert(false);
  }
  const timer = setTimeout(() => {
    setStatus(false);
  }, 3000);
};

 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="ระบบบันทึกการแจ้งซ่อมคอมพิวเตอร์" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button onClick={() => {CreateFix();}} variant="contained"  color="primary" startIcon={<AddCircleOutlinedIcon/>}> เพิ่มข้อมูลการแจ้งซ่อมสินค้า </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button style={{ marginLeft: 20 }} component={RouterLink} to="/Tablefix" variant="contained" startIcon={<CancelRoundedIcon/>}> ดูข้อมูล </Button>
     </Header>

     <Content>
       {status ? ( 
        <div>
          {alert ? ( 
              <Alert severity="success"> 
                <AlertTitle> บันทึกข้อมูลสำเร็จ </AlertTitle></Alert>) 
        : (     
          <Alert severity="warning"> 
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
              labelId="fixcomtypename-label"
              id="fixcomtypename"
              value={fixcomtype}
              onChange={FixcomtypehandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">เลือกประเภทคอมพิวเตอร์</InputLabel>

              {fixcomtypes.map((item: EntFixcomtype) => (
                <MenuItem value={item.id}>{item.fixcomtypename}</MenuItem>
              ))}
            </Select>

            
            <div className={classes.paper}><strong>แบรนด์</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              labelId="brandname-label"
              id="brandname"
              value={brand}
              onChange={BrandhandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">เลือกแบรนด์คอมพิวเตอร์</InputLabel>

              {brands.map((item: EntBrand) => (
                <MenuItem value={item.id}>{item.brandname}</MenuItem>
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
              id="productnumber"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={productnumber}
              onChange={handleProductnumberChange}
            />

            <div className={classes.paper}><strong>วันที่รับแจ้งซ่อม</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <ComputerTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="date"
              variant="standard"
              color="secondary"
              type="datetime-local"
              size="medium"
              value={date}
              onChange={handleDateChange}
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
              id="problemtype"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={problemtype}
              onChange={handleProblemtypeChange}
            />

            <div className={classes.paper}><strong>ลำดับคิว</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <ComputerTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="queue"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
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
