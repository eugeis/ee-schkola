import {Component, Input} from '@angular/core';
import {LibraryViewService} from '../../services/library-view.service}';

@Component({
  selector: 'app-library',
  templateUrl: './library-view.component.html',
  styleUrls: ['./library-view.component.scss'],
  providers: [LibraryViewService]
})

export class LibraryViewComponent {

    @Input() pageName = 'LibraryComponent';
       
    constructor(public libraryViewService: LibraryViewService) {}

}



