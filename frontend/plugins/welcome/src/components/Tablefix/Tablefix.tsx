import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import ComputerTwoToneIcon from '@material-ui/icons/ComputerTwoTone';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import { EntFix } from '../../api';

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
 
export default function Tablefix() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [fixs, setfixs] = useState<EntFix[]>([]);
  const [loading, setLoading] = useState(true);
  

  // get fixs
  useEffect(() => {
    const getFixs = async () => {
      const res:any = await http.listFix({ limit: 10, offset: 0 });
      setLoading(true);
      setfixs(res);
      console.log(res);
    };
    getFixs();
  }, [loading]);
  
  // delete button
  const deleteFixs = async (id: number) => {
    const res = await http.deleteFix({ id: id });
    setLoading(true);
  };

    // clear input form
    function clear() {
      setfixs([]);
    }
  
 
  // ui 
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`ระบบบันทึกการแจ้งซ่อมคอมพิวเตอร์`} type="ระบบแจ้งซ่อมคอมพิวเตอร์" >
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
    <Button variant="contained" color="secondary" href="/Fixpage" startIcon={<ComputerTwoToneIcon />}> เพิ่มข้อมูลบันทึกการแจ้งซ่อมสินค้า </Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/Group14" startIcon={<HomeRoundedIcon/>}> กลับหน้าหลัก </Button>
      </ContentHeader>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">ชื่อลูกค้า</TableCell>
           <TableCell align="center">ประเภทคอมพิวเตอร์</TableCell>
           <TableCell align="center">แบรนด์</TableCell>
           <TableCell align="center">หมายเลขผลิตภัณฑ์</TableCell>
           <TableCell align="center">วันที่รับแจ้งซ่อม</TableCell>
           <TableCell align="center">รายละเอียดการแจ้งซ่อม/ปัญา</TableCell>
           <TableCell align="center">ลำดับคิว</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {fixs.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.customer?.customername}</TableCell>
             <TableCell align="center">{item.edges?.fixcomtype?.fixcomtypename}</TableCell>
             <TableCell align="center">{item.edges?.brand?.brandname}</TableCell>
             <TableCell align="center">{item.productnumber}</TableCell>
             <TableCell align="center">{item.date}</TableCell>
             <TableCell align="center">{item.problemtype}</TableCell>
             <TableCell align="center">{item.queue}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteFixs(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
                 startIcon={<DeleteIcon/>}
                 href="/table"
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