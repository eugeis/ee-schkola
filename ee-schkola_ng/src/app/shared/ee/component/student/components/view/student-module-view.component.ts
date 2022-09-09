import {Component, Input} from '@angular/core';
import {StudentViewService} from '../../services/student-view.service}';

@Component({
  selector: 'app-student',
  templateUrl: './student-view.component.html',
  styleUrls: ['./student-view.component.scss'],
  providers: [StudentViewService]
})

export class StudentViewComponent {

    @Input() pageName = 'StudentComponent';
       
    constructor(public studentViewService: StudentViewService) {}

}



