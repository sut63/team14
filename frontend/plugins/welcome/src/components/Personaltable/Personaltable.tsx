//style
import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, ContentHeader, Header, Page, pageTheme } from '@backstage/core';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';

//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

//api
import { DefaultApi } from '../../api/apis';

//Entity
import { EntPersonal } from '../../api';

//icon
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import DeleteIcon from '@material-ui/icons/Delete';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  table: {
    minWidth: 650,
  },
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
  },
  }),
);
 
export default function Personaltable() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [loading, setLoading] = useState(true);

  const [personals, setPersonals] = useState<EntPersonal[]>([]);
  
  useEffect(() => {
    const getPersonals = async () => {
      const res:any = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(true);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();
  }, [loading]);
  
  // delete button
  const deletePersonals = async (id: number) => {
    const res = await http.deletePersonal({ id: id });
    setLoading(true);
  };

  // clear input form
  function clear() {
    setPersonals([]);
  }
  
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`ระบบข้อมูลบุคลากร`} type="ระบบแจ้งซ่อมคอมพิวเตอร์" >
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
      <ContentHeader title="ตารางแสดงผลข้อมูล">
      <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="secondary" href="/Personalcreate" startIcon={<PersonAddRoundedIcon />}> เพิ่มข้อมูลบุคลากร </Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/Personalwelcome" startIcon={<HomeRoundedIcon/>}> ย้อนกลับ </Button>
      </ContentHeader>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">คำนำหน้าชื่อ</TableCell>
           <TableCell align="center">ชื่อ-นามสกุล</TableCell>
           <TableCell align="center">เพศ</TableCell>
           <TableCell align="center">อีเมลล์</TableCell>
           <TableCell align="center">แผนก</TableCell>
           <TableCell align="center">จัดการ</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {personals.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.title?.titlename}</TableCell>
             <TableCell align="center">{item.personalname}</TableCell>
             <TableCell align="center">{item.edges?.gender?.gendername}</TableCell>
             <TableCell align="center">{item.email}</TableCell>
             <TableCell align="center">{item.edges?.department?.departmentname}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deletePersonals(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="primary"
                 startIcon={<DeleteIcon/>}
                 href="/Tablepersonal"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}