import {Attendance, AttendanceState, Course} from '@schkola/student/StudentApiBase'
import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {AttendanceDataService} from '../../service/attendance-data.service';

@Component({
  selector: 'app-attendance-form',
  templateUrl: './attendance-form.component.html',
  styleUrls: ['./attendance-form.component.scss']
})

export class AttendanceFormComponent implements OnInit {

    attendancestateEnum = this.attendanceDataService.loadEnumElement(AttendanceState);
    @Input() attendance: Attendance;

    constructor(public attendanceDataService: AttendanceDataService) {}

    ngOnInit(): void {
        if (this.attendance.student === undefined) {
            this.attendance.student = new Profile();
        }
        if (this.attendance.course === undefined) {
            this.attendance.course = new Course();
        }
    }
}




