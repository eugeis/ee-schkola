import {Attendance} from '@schkola/student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {AttendanceDataService} from '../../service/attendance-data.service';

@Component({
  selector: 'app-attendance-list',
  templateUrl: './attendance-entity-list.component.html',
  styleUrls: ['./attendance-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: AttendanceDataService}]
})

export class AttendanceListComponent implements OnInit {

    attendance: Attendance = new Attendance();

    tableHeader: Array<String> = [];

    constructor(public attendanceDataService: AttendanceDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'gender', 'first', 'last', 'birthName', 'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone', 'email', 'cellphone', 'photoData', 'photo', 'maritalState', 'childrenCount', 'first', 'last', 'church', 'member', 'services', 'name', 'level', 'id', 'other', 'profession', 'id', 'date', 'name', 'begin', 'end', 'first', 'last', 'name', 'start', 'end', 'dates', 'id', 'fee', 'description', 'id', 'hours', 'state', 'token', 'id'];
    }
}




