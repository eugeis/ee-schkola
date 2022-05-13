import {Component, Input, OnInit} from '@angular/core';
import {TemplateService} from '../../services/template.service';

@Component({
    selector: 'app-another-template',
    templateUrl: './child-template.component.html',
    styleUrls: ['./child-template.component.scss'],
})

export class ChildTemplateComponent implements OnInit {

    @Input() showElement: boolean;

    elementNameWithValue = [['firstname', 'string'], ['lastname', 'string']];
    show = true;

    constructor(public templateService: TemplateService) { }

    ngOnInit(): void {
        this.show = this.showElement;
        this.templateService.initElement(this.elementNameWithValue);
    }
}
