import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AcceptRequestPromptComponent } from './accept-request-prompt.component';

describe('AcceptRequestPromptComponent', () => {
  let component: AcceptRequestPromptComponent;
  let fixture: ComponentFixture<AcceptRequestPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AcceptRequestPromptComponent]
    });
    fixture = TestBed.createComponent(AcceptRequestPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
