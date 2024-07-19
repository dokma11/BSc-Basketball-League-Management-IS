import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddPlayerToListPromptComponent } from './add-player-to-list-prompt.component';

describe('AddPlayerToListPromptComponent', () => {
  let component: AddPlayerToListPromptComponent;
  let fixture: ComponentFixture<AddPlayerToListPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AddPlayerToListPromptComponent]
    });
    fixture = TestBed.createComponent(AddPlayerToListPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
