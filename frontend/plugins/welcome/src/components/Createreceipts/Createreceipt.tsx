//style
import React, {  useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, InputAdornment } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Alert, AlertTitle } from '@material-ui/lab';
import Swal from 'sweetalert2';
import Receipt from '.';

//api
import { DefaultApi } from '../../api/apis';

//Entity
import { EntPersonal, EntPaymentType, EntReceipt, EntAdminrepair, EntCustomer, EntProduct } from '../../api';

//icon
import NoteTwoToneIcon from '@material-ui/icons/NoteTwoTone';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import AddCircleTwoToneIcon from '@material-ui/icons/AddCircleTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';

//cookie
import { Cookies } from '../WelcomePage/Cookie';

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

export default function Personalpage() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);
  const [loading, setLoading] = useState(true);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  const [adminrepairs, setAdminrepairs] = React.useState<EntAdminrepair[]>([]);
  const [paymenttypes, setPaymenttypes] = React.useState<EntPaymentType[]>([]);
  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);
  const [products, setProducts] = React.useState<EntProduct[]>([]);

  const [personal, setPersonal] = useState(Number);
  const [adminrepair, setAdminrepair] = useState(Number);
  const [paymenttype, setPaymenttype] = useState(Number);
  const [customer, setCustomer] = useState(Number);
  const [product, setProduct] = useState(Number);

  const [addedTime, setAddedTime] = useState(String);
  const [serviceprovider, setServiceprovider] = useState(String);
  const [address, setAddress] = useState(String);
  const [productname, setProductname] = useState(String);

  useEffect(() => {
    const getCustomers = async () => {
      const res = await http.listCustomer({ limit: 10, offset: 0 });
      setLoading(false);
      setCustomers(res);
      console.log(res);
    };
    getCustomers();

    const getAdminrepairs = async () => {
      const res = await http.listAdminrepair({ limit: 10, offset: 0 });
      setLoading(false);
      setAdminrepairs(res);
      console.log(res);
    };
    getAdminrepairs();

    const getPaymenttypes = async () => {
      const res = await http.listPaymenttype({ limit: 10, offset: 0 });
      setLoading(false);
      setPaymenttypes(res);
      console.log(res);
    };
    getPaymenttypes();

    const getPersonals = async () => {
      const res = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();

  const getProducts = async () => {
    const res = await http.listProduct({ limit: 10, offset: 0 });
    setLoading(false);
    setProducts(res);
    console.log(res);
  };
  getProducts();

}, [loading]);

  const addedTimeChange = (event: any) => {
    setAddedTime(event.target.value as string);
  };
  
  const personalChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonal(event.target.value as number);
  };

  const paymenttypechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPaymenttype(event.target.value as number);
  };

  const adminrepairchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAdminrepair(event.target.value as number);
  };

  const customerchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCustomer(event.target.value as number);
  };

  const productchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProduct(event.target.value as number);
  };

  const handleServiceproviderChange = (event: any) => {
    setServiceprovider(event.target.value as string);
  };

  const handleAddressChange = (event: any) => {
    setAddress(event.target.value as string);
  };

  const handleProductnameChange = (event: any) => {
    setProductname(event.target.value as string);
  };
  let p = Number(personal)

  const receipt = {
    serviceprovider : serviceprovider,
    address : address,
    productname : productname,
    customer: customer,
    personal: Number(cookieID),
    adminrepair: adminrepair,
    paymentType: paymenttype,
    product: product,
    added : addedTime + ":00+07:00", 
  };

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'Serviceprovider':
        alertMessage("error","กรุณากรอกชื่อร้านค้าผู้ให้บริการ");
        return;
      case 'Address':
        alertMessage("error","กรุณากรอกที่อยู่ร้านค้า");
        return;
      case 'Productname':
        alertMessage("error","กรุณากรอกชื่อผลิตภัณฑ์");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  console.log(receipt)
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/receipts';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(receipt),
    };
  
    console.log(receipt); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
  
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
  
  /*const CreateReceipt = async () => {
    if ( (customer != null) && (personal != null) && (adminrepair != null) && (paymenttype != null) && (product != null) && (addedTime != "") &&(addedTime != null)){
    const receipt = {
      serviceprovider : serviceprovider,
      address : address,
      productname : productname,
      customer: customer,
      personal: Number(cookieID),
      adminrepair: adminrepair,
      paymentType: paymenttype,
      product: product,
      added : addedTime + ":00+07:00", 
    };

    console.log(receipt)
    const res: any = await http.createReceipt({ receipt: receipt });
      setStatus(true);
      if (res.id != '') {
        setAlert(true);
      } 
    }
      else {
        setStatus(true);
        setAlert(false);
      }
  };*/

return (
  <Page theme={pageTheme.tool}>
    <Header
      title="ระบบออกใบเสร็จ" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
      <Button 
        style={{ marginLeft: 20 }} 
        href="/"
        variant="contained"  
        startIcon={<ExitToAppIcon/>}
        > 
        ออกจากระบบ 
        </Button>
    </Header>
    <Content>
      <ContentHeader title="สร้างใบเสร็จ">
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button onClick={() => {save();}} variant="contained" color="primary" startIcon={<AddCircleOutlinedIcon/>}> Create new receipt </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button style={{ marginLeft: 20 }} component={RouterLink} to="/Tablereceipt" variant="contained" startIcon={<CancelRoundedIcon/>}>  ย้อนกลับ </Button>
      </ContentHeader>
  
      <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >

          <div className={classes.paper}><strong>ชื่อร้านค้าผู้ให้บริการ</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  
                </InputAdornment>
              ),
            }}
              id="serviceprovider"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={serviceprovider}
              onChange={handleServiceproviderChange}
            />

            <div className={classes.paper}><strong>รหัสลูกค้า</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="customerid"
              value={customer}
              onChange={customerchange}
            >
              {customers.map((item: EntCustomer) => (
                <MenuItem value={item.id}>{item.id}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>ชื่อ-นามสกุล</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="namecustomer"
              value={customer}
              onChange={customerchange}
            >
              {customers.map((item: EntCustomer) => (
                <MenuItem value={item.id}>{item.customername}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>เบอร์โทรศัพท์</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="phonenumber"
              value={customer}
              onChange={customerchange}
            >
              {customers.map((item: EntCustomer) => (
                <MenuItem value={item.id}>{item.phonenumber}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>ที่อยู่ร้านค้า</strong></div>
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

            <div className={classes.paper}><strong>ชื่อผลิตภัณฑ์</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  
                </InputAdornment>
              ),
            }}
              id="nameproduct"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={productname}
              onChange={handleProductnameChange}
            />

            <div className={classes.paper}><strong>รายละเอียดการซ่อม</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="adminrepair"
              value={adminrepair}
              onChange={adminrepairchange}
            >
              {adminrepairs.map((item: EntAdminrepair) => (
                <MenuItem value={item.id}>{item.equipmentdamate}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>ราคา</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="price"
              value={product}
              onChange={productchange}
            >
              {products.map((item: EntProduct) => (
                <MenuItem value={item.id}>{item.price}</MenuItem>
              ))}
            </Select>
            

            <div className={classes.paper}><strong>ประเภทการจ่ายเงิน</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="paymenttype"
              value={paymenttype}
              onChange={paymenttypechange}
            >
              <InputLabel className={classes.insideLabel}>เลือกประเภทการจ่ายเงิน</InputLabel>

              {paymenttypes.map((item: EntPaymentType) => (
                <MenuItem value={item.id}>{item.typename}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>วันที่รับแจ้งซ่อม</strong></div>
            <TextField 
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <AddCircleTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="date"
              variant="standard"
              color="secondary"
              type="date"
              size="medium"
              value={addedTime}
              onChange={addedTimeChange}
            />
            
          </FormControl>
        </form>
      </div>
    </Content>
   </Page>
 );
}