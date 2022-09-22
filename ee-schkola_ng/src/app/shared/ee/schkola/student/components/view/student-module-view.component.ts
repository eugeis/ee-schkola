import {Component, Input} from '@angular/core';
import {StudentViewService} from '../../service/student-module-view.service';

@Component({
  selector: 'app-student',
  templateUrl: './student-module-view.component.html',
  styleUrls: ['./student-module-view.component.scss'],
  providers: [StudentViewService]
})

export class StudentViewComponent {

    @Input() pageName = 'StudentComponent';
       
    constructor(public studentViewService: StudentViewService) {}

}



