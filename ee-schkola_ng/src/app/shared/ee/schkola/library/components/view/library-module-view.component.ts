import {Component, Input} from '@angular/core';
import {LibraryViewService} from '../../service/library-module-view.service';

@Component({
  selector: 'app-library',
  templateUrl: './library-module-view.component.html',
  styleUrls: ['./library-module-view.component.scss'],
  providers: [LibraryViewService]
})

export class LibraryViewComponent {

    @Input() pageName = 'LibraryComponent';
       
    constructor(public libraryViewService: LibraryViewService) {}

}



