import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InterviewInvitePromptComponent } from './interview-invite-prompt.component';

describe('InterviewInvitePromptComponent', () => {
  let component: InterviewInvitePromptComponent;
  let fixture: ComponentFixture<InterviewInvitePromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [InterviewInvitePromptComponent]
    });
    fixture = TestBed.createComponent(InterviewInvitePromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
