import {Attendance} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
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

    constructor(@Inject(AttendanceDataService) public attendanceDataService: AttendanceDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'student', 'date', 'course', 'hours', 'state', 'token', 'id'];
    }
}




