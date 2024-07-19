import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowRequestDetailsPromptComponent } from './show-request-details-prompt.component';

describe('ShowRequestDetailsPromptComponent', () => {
  let component: ShowRequestDetailsPromptComponent;
  let fixture: ComponentFixture<ShowRequestDetailsPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ShowRequestDetailsPromptComponent]
    });
    fixture = TestBed.createComponent(ShowRequestDetailsPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
